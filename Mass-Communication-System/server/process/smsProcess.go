package process2

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"Go-Projects/Mass-Communication-System/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProecss struct {
}

func (sp *SmsProecss) SendGroupSms(mes *message.Message) (err error) {

	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

func (sp *SmsProecss) SendMesToEachOnlineUser(data []byte, conn net.Conn) (err error) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
	return
}
