package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"log"
)

func main()  {

	// 服务器
	fmt.Println("Server have started......")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	for {

		// 接收客户端发送过来的数据
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}

		// 读取客户端发送过来的所有的数据
		request, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Receive a Message:%s\n",request)
	}


}