package main

import (
	"kongyixueyuan.com/publicChain/part3-Basic-Prototype/BLC"
	"fmt"
)

func main()  {

	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(genesisBlockchain)

	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])

}
