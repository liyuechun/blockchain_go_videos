package BLC

import "fmt"

// 先用它去查询余额
func (cli *CLI) getBalance(address string)  {

	fmt.Println("地址：" + address)

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}

	amount := utxoSet.GetBalance(address)

	fmt.Printf("%s一共有%d个Token\n",address,amount)

}
