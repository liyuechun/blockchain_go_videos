package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) printchain()  {

	if DBExists() == false {
		fmt.Println("数据不存在.......")
		os.Exit(1)
	}

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	blockchain.Printchain()

}