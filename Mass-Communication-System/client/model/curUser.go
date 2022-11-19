package model

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
