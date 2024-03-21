package resource

import (
	"encoding/json"
	"hertz_test/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/log"
	utils_redis "gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/redis"
	global "gitlab.appshahe.com/service-cloud-rec/sort_base.git/base_global"
)

var Gr = &GlobalResource{}

const (
	FidMapKey        = "sort_ad_with_game_feature_base"
	FidMapDefaultKey = "sort_ad_with_game_feature_base_default"
)

type GlobalResource struct {
	ServerConfig   *config.ServerConfig
	Env            string
	AdsRedisClient *AdsRedisClient
}

type AdsRedisClient struct {
	RedisN0                  *redis.Client
	RedisGroupCommonInstance *redis.Client
	RedisGroupModelInstance1 *redis.Client
	RedisGroupModelInstance2 *redis.Client
	RedisGroupUserInstance1  *redis.Client
	RedisGroupUserInstance2  *redis.Client
	RedisSspInstance         *redis.Client
	RedissonInstance         *redis.Client
}

type RedisConfig struct {
	Addr         string `json:"addr"`
	Password     string `json:"password"`
	DB           int    `json:"db"`
	DialTimeout  int64  `json:"dial_timeout"`
	ReadTimeout  int64  `json:"read_timeout"`
	WriteTimeout int64  `json:"write_timeout"`
	PoolSize     int    `json:"pool_size"`
	MinIdleConns int    `json:"min_idle_conns"`
}

func Init(env, configDir string) error {
	dir, dirErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if dirErr != nil {
		return dirErr
	}
	if env != "dev" {
		configDir = dir
	}

	Gr.Env = env
	Gr.ServerConfig = new(config.ServerConfig)
	panicIfErr(loadConfig(Gr.ServerConfig, "service", configDir, env))

	if adsRedisClient, err := initAdsRedisClient(configDir, env); err == nil {
		global.RedisClient = adsRedisClient.RedisN0
		Gr.AdsRedisClient = adsRedisClient
	} else {
		panicIfErr(err)
	}

	StartSchedule()

	log.Base().Info("init end")
	return nil
}

func loadConfig(config interface{}, name, configDir, env string) error {
	filePath := configDir
	if env != "" {
		filePath = filePath + "/" + name + "." + env
	}
	filePath = filePath + ".json"
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, config)
	if err != nil {
		return err
	}
	return nil
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initAdsRedisClient(configPath, env string) (*AdsRedisClient, error) {
	var redisConfigMap map[string]RedisConfig
	if err := loadConfig(&redisConfigMap, "redis", configPath, env); err != nil {
		return nil, err
	}
	adsRedisClient := &AdsRedisClient{}
	adsRedisClientReflect := reflect.ValueOf(adsRedisClient).Elem()
	for redisInstanceName, redisConfig := range redisConfigMap {
		if instance, err := getRedisClient(redisConfig); err == nil {
			adsRedisClientReflect.FieldByName(redisInstanceName).Set(reflect.ValueOf(instance))
		} else {
			return nil, err
		}
	}
	return adsRedisClient, nil
}

func getRedisClient(redisConfig RedisConfig) (*redis.Client, error) {
	RedisConfig := &redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		PoolSize:     redisConfig.PoolSize,
		DialTimeout:  time.Duration(redisConfig.DialTimeout),
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout),
		WriteTimeout: time.Duration(redisConfig.WriteTimeout),
		MinIdleConns: redisConfig.MinIdleConns,
	}
	return utils_redis.NewClient(RedisConfig)
}
