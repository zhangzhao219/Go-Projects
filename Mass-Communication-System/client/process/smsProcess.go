package process

import (
	"Go-Projects/Mass-Communication-System/client/utils"
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/json"
	"fmt"
)

type SmsProecss struct {
}

func (sp *SmsProecss) SendGroupSms(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
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
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
	return
}
