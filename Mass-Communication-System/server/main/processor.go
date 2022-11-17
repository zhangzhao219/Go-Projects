package main

import (
	"Go-Projects/Mass-Communication-System/common/message"
	process2 "Go-Projects/Mass-Communication-System/server/process"
	"Go-Projects/Mass-Communication-System/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送消息种类不同，决定调用哪个函数来实现
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录的逻辑
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (p *Processor) process2() (err error) {
	// 读取客户端发送的信息
	for {
		// 创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器退出")
				return err
			} else {
				fmt.Println("readPkg(conn) err", err)
			}
		}
		// fmt.Println(mes)
		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
