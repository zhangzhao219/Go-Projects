package config

import "github.com/go-redis/redis/v8"

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
