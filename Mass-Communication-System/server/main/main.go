package main

import (
	"Go-Projects/Mass-Communication-System/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	// 延时关闭连接
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通信协程有问题，err=", err)
		return
	}
}

// 完成对userdao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听.....")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	// 一旦监听成功，等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
