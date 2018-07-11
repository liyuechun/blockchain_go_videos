package BLC

import (
	"bytes"
	"log"
	"encoding/gob"
)

func handleVersion(request []byte,bc *Blockchain)  {

	var buff bytes.Buffer
	var payload Version

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	//Version
	//1. Version
	//2. BestHeight
	//3. 节点地址

	bestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if bestHeight > foreignerBestHeight {
		sendVersion(payload.AddrFrom,bc)
	} else if bestHeight < foreignerBestHeight {
		// 去向主节点要信息
		//sendGetBlocks(payload.AddrFrom)
	}


}

func handleAddr(request []byte,bc *Blockchain)  {

}

func handleGetblocks(request []byte,bc *Blockchain)  {

}

func handleGetData(request []byte,bc *Blockchain)  {

}

func handleBlock(request []byte,bc *Blockchain)  {

}

func handleTx(request []byte,bc *Blockchain)  {

}


func handleInv(request []byte,bc *Blockchain)  {

}