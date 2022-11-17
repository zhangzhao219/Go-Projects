package utils

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 将这些方法关联到结构体中
type Transfer struct {
	// 分析应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte // 传输时使用缓冲

}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("等待读取客户端发送的数据.....")
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	// fmt.Println("读到的长度为", buf[:4])
	// 转换为一个uint32类型
	var pkgLen = binary.BigEndian.Uint32(t.Buf[0:4])
	//  发送长度
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	// 把pkgLen反序列化成message
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度
	var pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	//  发送长度
	_, err = t.Conn.Write(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送data本身
	n, err := t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
