package process2

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"Go-Projects/Mass-Communication-System/server/model"
	"Go-Projects/Mass-Communication-System/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	// 表示Conn是哪个用户的
	UserId int
}

// 通知所有用户在线
func (u *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}

}

func (u *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
	}
}

// 编写serverProcessLogin函数，专门处理登录的请求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	fmt.Println(user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = err.Error()
		}

	} else {
		loginResMes.Code = 200
		// 因为用户登录成功，要将用户放入全局变量中以返回列表
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)
		u.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户的id放入到loginResMes.UsersIds
		for id := range userMgr.onlineUsers {
			loginResMes.UsersIds = append(loginResMes.UsersIds, id)
		}
		fmt.Println(user, "登录成功")
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
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 编写ServerProcessRegister函数，专门处理注册的请求
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 从mes中取出data，并反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal error, err=", err)
		return
	}
	// 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	// 声明一个RegisterResMes
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}

	} else {
		registerResMes.Code = 200
	}

	// 将loginResMes序列化
	data, err := json.Marshal(registerResMes)
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
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	return
}
