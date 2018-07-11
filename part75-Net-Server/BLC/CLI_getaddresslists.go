package BLC

import "fmt"

// 打印所有的钱包地址
func (cli *CLI) addressLists(nodeID string)  {

	fmt.Println("打印所有的钱包地址:")

	wallets,_ := NewWallets(nodeID)

	for address,_ := range wallets.WalletsMap {

		fmt.Println(address)
	}
}