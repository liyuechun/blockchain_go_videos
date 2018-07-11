package BLC

import "fmt"

func (cli *CLI) createWallet()  {

	wallets := NewWallets()

	wallets.CreateNewWallet()

	fmt.Println(wallets.Wallets)
}
