package BLC

import "fmt"

func (cli *CLI) TestMethod()  {


	fmt.Println("TestMethod")

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}

	utxoSet.ResetUTXOSet()

	//fmt.Println(blockchain.FindUTXOMap())
}
