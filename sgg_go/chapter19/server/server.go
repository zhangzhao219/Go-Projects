package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Println("服务器在等待客户端发送数据")
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端进行连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			return
		} else {
			fmt.Println("Accept success", conn, conn.RemoteAddr())
		}
		go process(conn)
	}
}
