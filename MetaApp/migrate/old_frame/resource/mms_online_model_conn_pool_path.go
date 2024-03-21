package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	global "gitlab.appshahe.com/service-cloud-rec/sort_base.git/base_global"
)

func AdsUpdateModelConnPool(ctx context.Context) error {
	return updateModelConnPool(ctx, global.RedisClient, "mms_online_model_conn_pool_path")
}

func updateModelConnPool(ctx context.Context, cli *redis.Client, key string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println("AdsUpdateModelConnPool execute")
	return nil
}
