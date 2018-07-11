package main

import "kongyixueyuan.com/publicChain/part24-persistence-cli/BLC"

func main()  {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	cli := BLC.CLI{blockchain}

	cli.Run()
}