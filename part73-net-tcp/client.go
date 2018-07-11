package main

import (
	"fmt"
	"io"
	"net"
	"bytes"
	"log"

)

func main() {

	sendMessage()
}

func sendMessage() {
	fmt.Println("客户端向服务器发送数据......")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic("error")
	}
	defer conn.Close()

	// 附带要发送的数据
	_, err = io.Copy(conn, bytes.NewReader([]byte("version")))
	if err != nil {
		log.Panic(err)
	}
}