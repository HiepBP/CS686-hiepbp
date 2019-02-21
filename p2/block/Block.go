package block

import (
	"p1"
)

//Header contains information of current block
type Header struct {
	height    int32
	timestamp int64 //UNIX timestamp 1550013938
	// hash_string = string(b.Header.Height) + string(b.Header.timestamp) + b.Header.ParentHash + b.Value.Root + string(b.Header.Size)
	// SHA3-256 encoded value of this string (follow this specific order)
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

//DecodeFromJSON function takes a string represents the json value of a block,
//decode the string back to a block instance
func DecodeFromJSON(json string) Block {
	var result Block
	return result
}

//EncodeToJSON function encode a block instance into json string
func (block *Block) EncodeToJSON() string {
	return ""
}
