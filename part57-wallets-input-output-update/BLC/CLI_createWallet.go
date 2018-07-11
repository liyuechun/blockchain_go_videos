package BLC

import "fmt"

func (cli *CLI) createWallet()  {

	wallets,_ := NewWallets()

	wallets.CreateNewWallet()

	fmt.Println(len(wallets.WalletsMap))
}
