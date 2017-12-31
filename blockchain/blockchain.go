package blockchain

import (
	"log"

	"github.com/spiermar/zerocoin/block"
)

var blockchain []*block.Block

// GetGenesisBlock returns the genesis block
func GetGenesisBlock() *block.Block {
	return blockchain[0]
}

// GetLatestBlock returns the last block on the blockchain
func GetLatestBlock() *block.Block {
	return blockchain[len(blockchain)-1]
}

// GenerateGenesisBlock generates the genesis block
func GenerateGenesisBlock(data string) {
	g := block.NewBlock(0, "", data)
	blockchain = append(blockchain, g)
}

// isValidGenesis validates the genesis block
func isValidGenesis(block *block.Block) bool {
	return GetGenesisBlock().Hash == block.Hash
}

// GenerateNextBlock generates a new block with the provided data
func GenerateNextBlock(data string) *block.Block {
	previousBlock := GetLatestBlock()
	nextIndex := previousBlock.Index + 1
	newBlock := block.NewBlock(nextIndex, previousBlock.Hash, data)
	blockchain = append(blockchain, newBlock)
	return newBlock
}

// isValidChain validates a blockchain
func isValidChain(blockchainToValidate []*block.Block) bool {
	if !isValidGenesis(blockchainToValidate[0]) {
		return false
	}
	for i, b := range blockchainToValidate {
		if i > 0 {
			if !block.IsValidBlock(b, blockchainToValidate[i-1]) {
				return false
			}
		}
	}
	return true
}

// replaceChain checks if a new chain is valid and longer and replaces the blockchain
func replaceChain(newBlocks []*block.Block) {
	if isValidChain(newBlocks) && len(newBlocks) > len(blockchain) {
		blockchain = newBlocks
	} else {
		log.Println("Invalid blockchain received.")
	}
}
