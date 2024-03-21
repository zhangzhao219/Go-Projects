package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	redisss "github.com/go-redis/redis/v8"
	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/redis"
)

// 声明一个全局的redisDb变量
var redisDbN0 *redisss.Client
var redisDbOnGC *redisss.Client
var redisDbPreGC *redisss.Client

// 根据redis配置初始化一个客户端
func initClientN0() (err error) {
	redisDbN0, err = redis.NewClient(&redisss.Options{
		Addr:     "r-2zeu8ajf31xvtwydzk.redis.rds.aliyuncs.com:6379", // redis地址
		Password: "JeeTSCKHudmXjc2zdLWt",                             // redis密码，没有则留空
		DB:       0,                                                  // 默认数据库，默认是0
	})
	return
}

func initClientOnlineGC() (err error) {
	redisDbOnGC, err = redis.NewClient(&redisss.Options{
		Addr:     "r-2zelmmx0a2mfmhy926.redis.rds.aliyuncs.com:6379", // redis地址
		Password: "UE52W79rxseB",                                     // redis密码，没有则留空
		DB:       8,                                                  // 默认数据库，默认是0
	})
	return
}

func initClientPreGC() (err error) {
	redisDbPreGC, err = redis.NewClient(&redisss.Options{
		Addr:     "r-2zelmmx0a2mfmhy926.redis.rds.aliyuncs.com:6379", // redis地址
		Password: "UE52W79rxseB",                                     // redis密码，没有则留空
		DB:       8,                                                  // 默认数据库，默认是0
	})
	return
}

type RecAdsConfigRequestParam struct {
	Aiid                int    `json:"aiid"`
	PackageName         string `json:"packageName"`
	Bid                 int    `json:"bid"`
	MaxNumOfImpressions int    `json:"maxNumOfImpressions"`
}

func read_packagename_bid() map[string]interface{} {
	content, err := ioutil.ReadFile("packagename_bid.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload["ad_bid"].(map[string]interface{})
}

func read_packagename_gameid() map[string]interface{} {
	content, err := ioutil.ReadFile("packagename_gameid.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload["ad_rec_list"].(map[string]interface{})
}

func main() {
	result := make(map[int64]*RecAdsConfigRequestParam)

	packagenameGameidMap := read_packagename_gameid()
	fmt.Println(len(packagenameGameidMap))
	for packagename, gameid := range packagenameGameidMap {
		result[int64(gameid.(float64))] = &RecAdsConfigRequestParam{
			PackageName: packagename,
		}
	}

	fmt.Println(len(result))

	packagenameBidMap := read_packagename_bid()
	fmt.Println(len(packagenameBidMap))
	for packagename, bid := range packagenameBidMap {
		for gameid := range result {
			if result[gameid].PackageName == packagename {
				result[gameid].Bid = int(bid.(float64))
			}
		}
	}

	for gameid, other := range result {
		fmt.Println(gameid, *other)
	}
	fmt.Println(len(result))

	err := initClientN0()
	fmt.Println(err)
	// initClientOnlineGC()
	// initClientPreGC()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println("fdsfds")
	n0map, _ := redisDbN0.HGetAll(ctx, "rec_sort:rec_ad_game_pkg_max_num_of_impressions:hash").Result()
	fmt.Println(n0map)

}

// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
// defer cancel()

// initClient()

// // var wg sync.WaitGroup
// // wg.Add(2)
// // starttime := time.Now().UnixNano()
// // go func() {
// // 	defer wg.Done()
// // 	redisDb.HSet(ctx, "rec_ads_config:gamerec_transfer_ad_config:gameid_aiid", "1", "2").Err()
// // }()
// // go func() {
// // 	defer wg.Done()
// // 	redisDb.HSet(ctx, "rec_ads_config:gamerec_and_recsort:gameid_other", "1", "2").Err()
// // }()
// // wg.Wait()
// // endtime := time.Now().UnixNano()
// starttime := time.Now().UnixNano()
// redisDb.HSet(ctx, "rec_ads_config:gamerec_transfer_ad_config:gameid_aiid", "1", "2").Err()
// redisDb.HSet(ctx, "rec_ads_config:gamerec_and_recsort:gameid_other", "1", "2").Err()
// endtime := time.Now().UnixNano()
// fmt.Println(endtime - starttime)

// func loadGameIdAiidFromRedis() map[int]string {
//
// 	defer cancel()
// 	gameRecTransferAdConfigMapResult := make(map[int]string, 0)

// 	redisGroupCommonInstance := redisDb
// 	gameRecTransferAdConfigMapMiddle, err := redis.HScanAll(ctx, redisGroupCommonInstance, "rec_ads_config:gameid_aiid:hash", 500)
// 	if err != nil {
// 		log.Base().Errorf("redis.HScanAll failed, err : %v", err)
// 		return gameRecTransferAdConfigMapResult
// 	}
// 	for gameid, aiid := range gameRecTransferAdConfigMapMiddle {
// 		gameidInt, parseIntErr := strconv.ParseInt(gameid, 10, 64)
// 		if parseIntErr != nil {
// 			log.Base().Errorf("gameid : %v, gameidInt : %v, parseInt err : %v", gameid, gameidInt, parseIntErr)
// 			continue
// 		}
// 		gameRecTransferAdConfigMapResult[int(gameidInt)] = aiid
// 	}
// 	return gameRecTransferAdConfigMapResult
// }

// func main2() {
// 	err := initClient()
// 	if err != nil {
// 		//redis连接错误
// 		panic(err)
// 	}
// 	fmt.Println("Redis连接成功")

// 	redisDb.Set("key", "11", 0)

// 	// 定义一个回调函数，用于处理事务逻辑
// 	fn := func(tx *redis.Tx) error {
// 		// 先查询下当前watch监听的key的值
// 		v, err := tx.Get("key").Result()
// 		if err != nil && err != redis.Nil {
// 			return err
// 		}

// 		// 这里可以处理业务
// 		fmt.Println(v)

// 		// 如果key的值没有改变的话，Pipelined函数才会调用成功
// 		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
// 			// 在这里给key设置最新值
// 			pipe.Set("key", "new value", 0)
// 			return nil
// 		})
// 		return err
// 	}

// 	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
// 	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
// 	redisDb.Watch(fn, "key")

// }
