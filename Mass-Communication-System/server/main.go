package main

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*4)
	fmt.Println("等待读取客户端发送的数据.....")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	// fmt.Println("读到的长度为", buf[:4])
	// 转换为一个uint32类型
	var pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//  发送长度
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	// 把pkgLen反序列化成message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func process(conn net.Conn) {
	// 延时关闭连接
	defer conn.Close()

	// 读取客户端发送的信息
	for {
		// 将读取数据包直接封装成一个函数
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器退出")
				return
			} else {
				fmt.Println("readPkg(conn) err", err)
			}
		}
		fmt.Println(mes)
	}

}

func main() {
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
