package BLC


func (cli *CLI) resetUTXOSet()  {

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}

	utxoSet.ResetUTXOSet()

}
