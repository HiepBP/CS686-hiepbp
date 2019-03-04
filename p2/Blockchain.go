package p2

import (
	"encoding/json"
	"fmt"
)

//BlockChain contain a map(which maps to block height to a list of blocks)
//and length equals to the highest block height
type BlockChain struct {
	chain  map[int32][]Block
	length int32
}

type BlockChainJson struct {
	chain []BlockJson
}

type BlockJson struct {
	Elements map[string]string `json:" "`
}

func NewBlockChain() *BlockChain {
	chain := make(map[int32][]Block)
	length := int32(0)
	return &BlockChain{chain, length}
}

func (blockChain *BlockChain) Length() int32 {
	return blockChain.length
}

//Get function take a height then return the lists of blocks stored in that height
//or None if the height doesn't exist
func (blockChain *BlockChain) Get(height int32) []Block {
	//Check if height is valid
	if height > blockChain.length || height <= 0 {
		return nil
	}
	return blockChain.chain[height]
}

//Insert function takes a block then insert it into BlockChain.chain
//If the list has already contained that block's hash,
//ignore it because we don't store duplicate blocks; if not, insert the block into the list.
//QUESTION: do we need to check parentHash in previous height before Insert
//QUESTION: if block chain at height 4, can we insert a new block at height 1 or 2?
func (blockChain *BlockChain) Insert(newBlock Block) {
	//Check if height is valid
	if newBlock.Header.Height <= 0 || newBlock.Header.Height-blockChain.length > 1 {
		return
	}
	if listBlock, ok := blockChain.chain[newBlock.Header.Height]; ok { //Already has that height
		if !containHash(listBlock, newBlock) {
			listBlock = append(listBlock, newBlock)
			blockChain.chain[newBlock.Header.Height] = listBlock
		}
	} else { //New height
		blockChain.chain[newBlock.Header.Height] = []Block{newBlock}
		blockChain.length = newBlock.Header.Height
	}
}

//EncodeToJSON function encode the blockchain instance into a json string
func (blockChain *BlockChain) EncodeToJSON() (string, error) {
	var result string
	var elements []map[string]interface{}
	var element map[string]interface{}
	for _, blockHeight := range blockChain.chain {
		for _, block := range blockHeight {
			element = make(map[string]interface{})
			jsonBlock, err := block.EncodeToJSON()
			if err != nil {
				return result, err
			}
			//Pass jsonBlock back into map of key value in json
			err = json.Unmarshal([]byte(jsonBlock), &element)
			if err != nil {
				return result, err
			}
			elements = append(elements, element)
		}
	}
	jsonByte, _ := json.Marshal(elements)
	result = string(jsonByte)
	fmt.Println(string(jsonByte))
	return result, nil
}

//DecodeFromJSON function call by blockchain instance and take json string
//decode that string back to blockchain instance and copy everthing to the current blockchain
func DecodeJsonToBlockChain(jsonString string) (*BlockChain, error) {
	result := NewBlockChain()
	var elements []map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &elements)
	for _, element := range elements {
		jsonElement, _ := json.Marshal(element)
		block, _ := DecodeFromJSON(string(jsonElement))
		result.Insert(block)
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func containHash(listBlock []Block, block Block) bool {
	for _, currBlock := range listBlock {
		if currBlock.Header.Hash == block.Header.Hash {
			return true
		}
	}
	return false
}
