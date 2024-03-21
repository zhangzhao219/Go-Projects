package main

import (
	user_feature "kitextest/kitex_gen/user_feature/userfeaturerpcservice"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

const userFeatureDatawarehouseName = "user-feature-datawarehouse"

func main() {

	opts := []server.Option{
		server.WithServiceAddr(
			&net.TCPAddr{
				IP:   net.ParseIP("0.0.0.0"),
				Port: 10001,
			},
		),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: userFeatureDatawarehouseName,
			},
		),
		server.WithMaxConnIdleTime(
			time.Second * 15,
		),
	}

	svr := user_feature.NewServer(new(UserFeatureRPCServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
