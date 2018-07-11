package BLC

import "fmt"

type Wallets struct {
	Wallets map[string]*Wallet
}



// 创建钱包集合
func NewWallets() *Wallets {

	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return wallets
}

// 创建一个新钱包
func (w *Wallets) CreateNewWallet()  {

	wallet := NewWallet()
	fmt.Printf("Address：%s\n",wallet.GetAddress())
	w.Wallets[string(wallet.GetAddress())] = wallet
}


