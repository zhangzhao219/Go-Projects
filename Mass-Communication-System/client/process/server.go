package process

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"Go-Projects/Mass-Communication-System/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 显示成功后的界面
func ShowMenu() {
	fmt.Println("---------------恭喜登录成功---------------")
	fmt.Println("------------1 显示在线用户列表")
	fmt.Println("---------------2 发送消息")
	fmt.Println("---------------3 信息列表")
	fmt.Println("---------------4 退出系统")
	fmt.Println("请选择(1-4): ")
	var key int
	var content string
	smsProcess := &SmsProecss{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		// fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你想对大家说的话")
		fmt.Scanln(&content)
		smsProcess.SendGroupSms(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统......")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}

// 和服务器端保持通讯
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		// 如果读取到消息,下一步进行处理
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal err=", err)
			}
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回未知消息")
		}
	}
}
