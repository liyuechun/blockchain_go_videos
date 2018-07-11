package main

import (
	"kongyixueyuan.com/publicChain/part53-wallets/BLC"
	"fmt"
)

func main() {

	wallets := BLC.NewWallets()

	fmt.Println(wallets.Wallets)

	wallets.CreateNewWallet()
	wallets.CreateNewWallet()

	fmt.Println(wallets.Wallets)

}