package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}
	fmt.Println("conn success", conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err", err)
		}
		line = strings.Trim(line, "\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			return
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Println("bytes:", n)
	}

}
