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
		fmt.Println("登录成功")
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
