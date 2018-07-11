package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"crypto/rand"
)

//big.Int

type Wallet struct {
	//1. 私钥
	PrivateKey ecdsa.PrivateKey

	//2. 公钥
	PublicKey  []byte
}


// 创建钱包
func NewWallet() *Wallet {

	privateKey,publicKey := newKeyPair()

	return &Wallet{privateKey,publicKey}
}


// 通过私钥产生公钥
func newKeyPair() (ecdsa.PrivateKey,[]byte) {

	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}