package BLC


// 转账
func (cli *CLI) send(from []string,to []string,amount []string,nodeID string)  {


	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from,to,amount,nodeID)

	utxoSet := &UTXOSet{blockchain}

	//转账成功以后，需要更新一下
	utxoSet.Update()

}

