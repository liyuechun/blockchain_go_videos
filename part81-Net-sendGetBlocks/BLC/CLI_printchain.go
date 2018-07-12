package BLC


func (cli *CLI) printchain(nodeID string)  {

	blockchain := BlockchainObject(nodeID)

	defer blockchain.DB.Close()

	blockchain.Printchain()

}