package BLC

type Blockchain struct {
	Blocks []*Block  // 存储有序的区块
}





//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	// 创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.......")
	// 返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}