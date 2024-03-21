package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"temp/proto/user_feature_datawarehouse/user_feature"

	utils_redis "gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/redis"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
)

const (
	fidCacheRedisKey   = "user_feature_warehouse:common:fid:cache:%s"
	fidCacheExpireTime = 10
)

type RecentResult struct {
	ClickList      []string
	ClickTimeList  []string
	PlayList       []string
	PlayTimeList   []string
	PlayListT1     []string
	PlayTimeListT1 []string
	PlayIdx        map[string]int64
}

func main() {

	redisConf := &redis.Options{
		Addr:     "121.37.98.68:6379",
		Password: "3jJEIoWm3jJEIoWmYPw3tT7Ndtmy8pCOmZhBTSKRCTvswzrvUAg",
		DB:       0,
	}
	rediscli, err := utils_redis.NewClient(redisConf)
	if err != nil {
		log.Fatal("init redis fail")
		return
	}

	recentResult := &RecentResult{
		ClickList:      []string{"ClickList1", "ClickList2", "ClickList3"},
		ClickTimeList:  []string{"ClickTimeList1", "ClickTimeList2"},
		PlayList:       []string{"PlayList1", "PlayList2", "PlayList3"},
		PlayTimeList:   []string{"PlayTimeList1", "PlayTimeList2"},
		PlayListT1:     []string{"PlayListT11", "PlayListT12"},
		PlayTimeListT1: []string{"PlayTimeListT11", "PlayTimeListT12", "PlayTimeListT13"},
		PlayIdx:        make(map[string]int64),
	}
	recentResult.PlayIdx["sdf"] = 12

	mode := 1

	var recentResultMap map[string]interface{}

	starttime := time.Now().UnixNano()

	if mode == 1 {
		recentResultMap, err = structToMap(*recentResult)
		if err != nil {
			log.Fatal("structToMap fail")
		}

		WriteFidBytes2Redis(rediscli, "gfd", recentResult)
	} else {
		recentResultMap, err = structToMap(*recentResult)
		if err != nil {
			log.Fatal("structToMap fail")
		}
		WriteMap2Redis(rediscli, "gfd", recentResult)
	}

	middle1time := time.Now().UnixNano()
	fmt.Println("store time:", middle1time-starttime)
	fmt.Println(recentResultMap)
	middle2time := time.Now().UnixNano()

	ctx, cancel := context.WithTimeout(context.Background(), 5*100*time.Millisecond)
	if mode == 1 {
		defer cancel()
		recentResultMap, _ = getUserFeatureFromRedis(ctx, "gfd", rediscli)
	} else {
		recentResultMap, _ = getUserFeatureFromRedis2(ctx, "gfd", rediscli)
	}
	endtime := time.Now().UnixNano()
	fmt.Println("extract time:", endtime-middle2time)
	fmt.Println(recentResultMap)
}

func structToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; get %T", v)
	}
	t := reflect.TypeOf(in)
	for i := 0; i < v.NumField(); i++ {
		out[t.Field(i).Name] = v.Field(i).Interface()
	}
	return out, nil
}

func WriteFidBytes2Redis(cli *redis.Client, uuid string, fidRecentResult *RecentResult) {
	fidProtoByte, compressErr := compressFid(fidRecentResult)
	if compressErr != nil {
		log.Fatalf("compress fid err : %v", compressErr)
		return
	}
	redisKey := fmt.Sprintf(fidCacheRedisKey, uuid)
	newCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, err := cli.Set(newCtx, redisKey, fidProtoByte, time.Duration(fidCacheExpireTime)*time.Minute).Result()
	if err != nil {
		log.Fatalf("write fidBytes 2 redis : %v", err)
	}
}

func compressFid(fidRecentResult *RecentResult) ([]byte, error) {
	userFeatureCache := &user_feature.UserFeatureCache{
		ClickList:      fidRecentResult.ClickList,
		ClickTimeList:  fidRecentResult.ClickTimeList,
		PlayList:       fidRecentResult.PlayList,
		PlayTimeList:   fidRecentResult.PlayTimeList,
		PlayListT1:     fidRecentResult.PlayListT1,
		PlayTimeListT1: fidRecentResult.PlayTimeListT1,
		PlayIdx:        fidRecentResult.PlayIdx,
	}
	protoByte, protoMarshalErr := proto.Marshal(userFeatureCache)
	if protoMarshalErr != nil {
		return nil, protoMarshalErr
	}
	return protoByte, nil
}

func getUserFeatureFromRedis(ctx context.Context, uuid string, cli *redis.Client) (map[string]interface{}, error) {
	newCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	redisKey := fmt.Sprintf(fidCacheRedisKey, uuid)
	str, redisErr := cli.Get(newCtx, redisKey).Result()
	if redisErr != nil {
		return make(map[string]interface{}), redisErr
	}
	userFeatureCache := &user_feature.UserFeatureCache{}
	protoErr := proto.Unmarshal([]byte(str), userFeatureCache)
	if protoErr != nil {
		log.Fatalf("proto unmarshal user feature cache err : %v", protoErr)
		return make(map[string]interface{}), protoErr
	}
	recentResultMap, structToMapErr := structToMap(*userFeatureCache)
	if structToMapErr != nil {
		return make(map[string]interface{}), structToMapErr
	}
	return recentResultMap, nil
}

func WriteMap2Redis(cli *redis.Client, uuid string, fidRecentResult *RecentResult) {
	fidProtoByte, compressErr := compressFid(fidRecentResult)
	if compressErr != nil {
		log.Fatalf("compress fid err : %v", compressErr)
		return
	}
	redisKey := fmt.Sprintf(fidCacheRedisKey, uuid)
	newCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, err := cli.HSet(newCtx, redisKey, fidProtoByte).Result()
	if err != nil {
		log.Fatalf("write fidBytes 2 redis : %v", err)
	}
}

func getUserFeatureFromRedis2(ctx context.Context, uuid string, cli *redis.Client) (map[string]interface{}, error) {
	newCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	redisKey := fmt.Sprintf(fidCacheRedisKey, uuid)
	str, redisErr := cli.HGet(newCtx, redisKey, redisKey).Result()
	if redisErr != nil {
		return make(map[string]interface{}), redisErr
	}
	userFeatureCache := &user_feature.UserFeatureCache{}
	protoErr := proto.Unmarshal([]byte(str), userFeatureCache)
	if protoErr != nil {
		log.Fatalf("proto unmarshal user feature cache err : %v", protoErr)
		return make(map[string]interface{}), protoErr
	}
	recentResultMap, structToMapErr := structToMap(*userFeatureCache)
	if structToMapErr != nil {
		return make(map[string]interface{}), structToMapErr
	}
	return recentResultMap, nil
}
