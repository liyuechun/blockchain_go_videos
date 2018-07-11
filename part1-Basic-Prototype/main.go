package main

import (
	"kongyixueyuan.com/publicChain/part1-Basic-Prototype/BLC"
	"fmt"
)

func main()  {

	block := BLC.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)

}
