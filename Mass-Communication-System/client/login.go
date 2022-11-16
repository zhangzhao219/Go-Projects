package main

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// fmt.Printf("userId=%d, userPed=%s\n", userId, userPwd)
	// return nil
	// 连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	// 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	mes.Data = string(data)
	// 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	// data为发送的消息
	// 先把data的长度发送给服务器
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//  发送长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	// fmt.Println("客户端发送的消息长度为", len(data))
	// fmt.Println("客户端发送的消息内容为", string(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	// 处理服务器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) error, err=", err)
		return
	}
	// 将mes的data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
