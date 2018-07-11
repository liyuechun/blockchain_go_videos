package BLC

import (
	"fmt"
	"io"
	"bytes"
	"log"
	"net"
)

func sendVersion(toAddress string,bc *Blockchain)  {


	bestHeight := bc.GetBestHeight()
	payload := gobEncode(Version{NODE_VERSION, bestHeight, nodeAddress})

	request := append(commandToBytes(COMMAND_VERSION), payload...)

	sendData(toAddress,request)


}



func sendData(to string,data []byte)  {

	fmt.Println("客户端向服务器发送数据......")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		panic("error")
	}
	defer conn.Close()

	// 附带要发送的数据
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}