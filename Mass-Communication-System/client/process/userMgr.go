package process

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"fmt"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 在客户端显示当前在线的用户

func outputOnlineUser() {
	fmt.Println("当前在线用户列表")
	for id, user := range onlineUsers {
		fmt.Println(id, user)
	}
}

// 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
