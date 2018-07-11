package main

import (
	"kongyixueyuan.com/publicChain/part19-persistence-Iterator/BLC"
)

func main()  {

	// 创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	 //新区块
	blockchain.AddBlockToBlockchain("Send 100RMB To zhangqiang")

	blockchain.AddBlockToBlockchain("Send 200RMB To changjingkong")

	blockchain.AddBlockToBlockchain("Send 300RMB To juncheng")

	blockchain.AddBlockToBlockchain("Send 50RMB To haolin")

	blockchain.Printchain()

}
