package main

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
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

func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//  发送长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送data本身
	n, err := conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
