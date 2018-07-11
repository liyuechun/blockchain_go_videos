package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {

	msg := "http://liyuechun.org"
	encoded := base64.URLEncoding.EncodeToString([]byte(msg))

	fmt.Println(encoded) //SGVsbG8sIOS4lueVjA==

	//aHR0cDovL2xpeXVlY2h1bi5vcmc
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))


	



}
