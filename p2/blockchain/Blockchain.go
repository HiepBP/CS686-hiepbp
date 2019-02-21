package blockchain

import (
	"../block"
)

//BlockChain contain a map(which maps to block height to a list of blocks)
//and length equals to the highest block height
type BlockChain struct {
	chain  map[int32][]block.Block
	length int32
}

//Get function take a height then return the lists of blocks stored in that height
//or None if the height doesn't exist
func (blockChain *BlockChain) Get(height int32) []block.Block {
	var result []block.Block
	return result
}

//Insert function takes a block then insert it into BlockChain.chain
func (blockChain *BlockChain) Insert(block block.Block) {

}

//EncodeToJSON function encode the blockchain instance into a json string
func (blockChain *BlockChain) EncodeToJSON() string {
	var result string
	return result
}

//DecodeFromJSON function call by blockchain instance and take json string
//decode that string back to blockchain instance and copy everthing to the current blockchain
func (blockChain *BlockChain) DecodeFromJSON(json string) {

}
