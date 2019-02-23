package p2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"../p1"
	"golang.org/x/crypto/sha3"
)

//Header contains information of current block
type Header struct {
	Height    int32 `json:"height"`
	Timestamp int64 `json:"timestamp"` //UNIX timestamp 1550013938
	// hash_string = string(b.Header.Height) + string(b.Header.timestamp) + b.Header.ParentHash + b.Value.Root + string(b.Header.Size)
	// SHA3-256 encoded value of this string (follow this specific order)
	Hash       string
	ParentHash string `json:"parentHash"`
	Size       int32  `json:"size"`
}

//Block contains header and value
type Block struct {
	Header Header
	Value  p1.MerklePatriciaTrie `json:"mpt"`
}

//BlockCustom is the json model for Block
type BlockCustom struct {
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timeStamp"` //UNIX timestamp 1550013938
	Height    int32  `json:"height"`
	// hash_string = string(b.Header.Height) + string(b.Header.timestamp) + b.Header.ParentHash + b.Value.Root + string(b.Header.Size)
	// SHA3-256 encoded value of this string (follow this specific order)
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	Value      map[string]string `json:"mpt"`
}

//Initial function takes arguments(such as height, parentHash, and value of MPT type)
//Then forms a block
func Initial(height int32, timestamp int64, parentHash string, mpt p1.MerklePatriciaTrie) Block {
	var result Block
	var size int32

	//Init value
	size = int32(len([]byte(fmt.Sprintf("%v", mpt))))
	stringMix := fmt.Sprintf("%v", height) +
		fmt.Sprintf("%v", timestamp) +
		parentHash +
		mpt.Root() +
		fmt.Sprintf("%v", size)
	hash := sha3.Sum256([]byte(stringMix))

	//Assign value
	result.Value = mpt
	result.Header.Hash = hex.EncodeToString(hash[:])
	result.Header.Height = height
	result.Header.ParentHash = parentHash
	result.Header.Size = size
	result.Header.Timestamp = timestamp
	return result
}

//DecodeFromJSON function takes a string represents the json value of a block,
//decode the string back to a block instance
func DecodeFromJSON(jsonString string) (Block, error) {
	// result := &BlockCustom{
	// 	Header: Header{},
	// 	Value:  {},
	// }

	var result Block
	var blockCustom BlockCustom

	err := json.Unmarshal([]byte(jsonString), &blockCustom)
	if err != nil {
		return result, err
	}
	result = blockCustom.toBlock()
	return result, nil
}

//EncodeToJSON function encode a block instance into json string
func (block *Block) EncodeToJSON() (string, error) {
	var result string

	blockCustom := block.toBlockCustom()

	blockByte, err := json.Marshal(&blockCustom)
	if err != nil {
		return result, err
	}
	result = string(blockByte)
	return result, nil
}

//This function convert from BlockCustom to Block
func (block *BlockCustom) toBlock() Block {
	var result Block

	//Assign value
	result.Value = *p1.InitMPT(block.Value)
	result.Header.Hash = block.Hash
	result.Header.Height = block.Height
	result.Header.ParentHash = block.ParentHash
	result.Header.Size = block.Size
	result.Header.Timestamp = block.Timestamp

	return result
}

//This function convert from Block to BlockCustom
func (block *Block) toBlockCustom() BlockCustom {
	var result BlockCustom

	//Assign value
	result.Value = block.Value.GetListKeyValue()
	result.Hash = block.Header.Hash
	result.Height = block.Header.Height
	result.ParentHash = block.Header.ParentHash
	result.Size = block.Header.Size
	result.Timestamp = block.Header.Timestamp

	return result
}
