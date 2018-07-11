package main

import (
	"kongyixueyuan.com/publicChain/part9-Serialize-DeserializeBlock/BLC"
	"fmt"
)

func main()  {

	//data string,height int64,prevBlockHash []byte
	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)


	bytes := block.Serialize()

	fmt.Println(bytes)

	block = BLC.DeserializeBlock(bytes)

	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)
}
