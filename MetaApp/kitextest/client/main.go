package main

import (
	"context"
	"kitextest/kitex_gen/user_feature"
	"kitextest/kitex_gen/user_feature/userfeaturerpcservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

const userFeatureDatawarehouseName = "user-feature-datawarehouse"

func main() {
	opts := []client.Option{
		client.WithHostPorts("0.0.0.0:10001"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithLongConnection(
			connpool.IdleConfig{
				MaxIdlePerAddress: 50,
				MaxIdleGlobal:     500,
				MaxIdleTimeout:    time.Second * 5,
				MinIdlePerAddress: 2,
			},
		),
	}
	clientRPC, err := userfeaturerpcservice.NewClient(userFeatureDatawarehouseName, opts...)
	if err != nil {
		klog.Errorf("new client failed, error=%v", err)
	}

	for {
		for i := 0; i < 50; i++ {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
				defer cancel()
				resp, err := clientRPC.GetUserFeature(ctx, &user_feature.Request{})
				if err != nil {
					klog.Errorf("connect failed, error=%v", err)
				}
				klog.Infof("resp %v", resp)
				time.Sleep(time.Second)
			}()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
