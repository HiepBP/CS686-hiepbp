package block

import (
	"p1"
)

//Header contains information of current block
type Header struct {
	height     int32
	timestamp  int64 //UNIX timestamp 1550013938
	hash       string
	parentHash string
	size       int32
}

//Block contains header and value
type Block struct {
	header Header
	value  p1.MerklePatriciaTrie
}

//Initial function takes arguments(such as height, parentHash, and value of MPT type)
//Then forms a block
func Initial() {

}

//DecodeFromBase64 function takes a string represents the base64 hashed value of a block,
//decode the string back to a block instance
func DecodeFromBase64(base64String string) Block {
	var result Block
	return result
}

//EncodeToBase64 function encode a block instance into base64 format
func (block *Block) EncodeToBase64() string {
	return ""
}
