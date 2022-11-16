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

// 编写serverProcessLogin函数，专门处理登录的请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 从mes中取出data，并反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error, err=", err)
		return
	}
	// 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 声明一个LoginResMes
	var loginResMes message.LoginResMes
	// 如果用户的id为100，密码为123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}
	// 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal error, err=", err)
		return
	}
	// 将data赋值给resMes
	resMes.Data = string(data)
	// 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error, err=", err)
		return
	}
	// 发送data，封装到writePkg函数
	err = writePkg(conn, data)
	return
}

// 根据客户端发送消息种类不同，决定调用哪个函数来实现
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录的逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
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
		// fmt.Println(mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
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
