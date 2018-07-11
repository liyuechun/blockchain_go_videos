package BLC

import (
	//"fmt"
	//"encoding/hex"
)

func (cli *CLI) TestMethod()  {


	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	utxo_set := &UTXOSet{blockchain}
	utxo_set.ResetUTXOSet()

}
