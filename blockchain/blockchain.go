package blockchain

import (
	"log"

	"github.com/spiermar/zerocoin/block"
	"github.com/spiermar/zerocoin/proto"
)

var blockchain []*proto.Block

// GetGenesisBlock returns the genesis block
func GetGenesisBlock() *proto.Block {
	return blockchain[0]
}

// GetLatestBlock returns the last block on the blockchain
func GetLatestBlock() *proto.Block {
	return blockchain[len(blockchain)-1]
}

// GenerateGenesisBlock generates the genesis block
func GenerateGenesisBlock(data string) {
	g := block.NewBlock(0, "", data)
	blockchain = append(blockchain, g)
}

// isValidGenesis validates the genesis block
func isValidGenesis(block *proto.Block) bool {
	return GetGenesisBlock().Hash == block.Hash
}

// GenerateNextBlock generates a new block with the provided data
func GenerateNextBlock(data string) *proto.Block {
	previousBlock := GetLatestBlock()
	nextIndex := previousBlock.Index + 1
	newBlock := block.NewBlock(nextIndex, previousBlock.Hash, data)
	blockchain = append(blockchain, newBlock)
	return newBlock
}

// isValidChain validates a blockchain
func isValidChain(blockchainToValidate []*proto.Block) bool {
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

// ReplaceChain checks if a new chain is valid and longer and replaces the blockchain
func ReplaceChain(newBlocks []*proto.Block) bool {
	if isValidChain(newBlocks) && len(newBlocks) > len(blockchain) {
		blockchain = newBlocks
	} else {
		log.Println("Invalid blockchain received.")
		return false
	}
	return true
}

// AddBlockToChain checks the block and adds it to the chain
func AddBlockToChain(newBlock *proto.Block) bool {
	latestBlockHeld := GetLatestBlock()
	if block.IsValidBlock(newBlock, latestBlockHeld) {
		blockchain = append(blockchain, newBlock)
	} else {
		log.Println("Invalid block received.")
		return false
	}
	return true
}
