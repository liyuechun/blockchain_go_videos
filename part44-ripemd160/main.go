package main

import (
	"golang.org/x/crypto/ripemd160"
	"fmt"
)

func main()  {

	//sha256 256bit
	//ripemd160 160bit


	// 160
	// b66140b4bfd22da44399f352b07182864098123f
	// bit 160 20个字节
	hasher := ripemd160.New()

	hasher.Write([]byte("http://liyuechun.org"))

	bytes := hasher.Sum(nil)

	fmt.Printf("%x\n",bytes)

}
