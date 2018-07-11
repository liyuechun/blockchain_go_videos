package main

import (
	"kongyixueyuan.com/publicChain/part51-wallet-address/BLC"
	"fmt"
)

func main() {

	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("address：%s\n",address)
//	1EESFLjfwLqoSq4LJA8eymoPi5pc5Sq6VC
//  1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa


}