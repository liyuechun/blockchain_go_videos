package main

import (
	"kongyixueyuan.com/publicChain/part52-wallet-address/BLC"
	"fmt"
)

func main() {

	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("address：%s\n",address)

	//	1Gh5pL2uFsS8AFuSkd6za521FzsWHJcXDb
	//	1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

	isValid := BLC.IsValidForAdress(address)

	fmt.Printf("%s 这个地址为 %v\n",address,isValid)

}