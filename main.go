package main

import (
	"fmt"

	"./p1"
	"./p2"
)

func main() {
	mpt := p1.MerklePatriciaTrie{}
	mpt.Initial()
	mpt.Insert("hello", "world")
	mpt.Insert("charles", "ge")
	b1 := p2.Initial(1, 1234567890, "genesis", mpt)
	b2 := p2.Initial(2, 1234567890, b1.Header.Hash, mpt)
	bc := p2.NewBlockChain()
	bc.Insert(b1)
	bc.Insert(b2)
	jsonB1, _ := b1.EncodeToJSON()
	jsonB2, _ := b2.EncodeToJSON()
	fmt.Println(jsonB1)
	fmt.Println(jsonB2)
	fmt.Println(bc.Length())

	// json := "{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}"
	// b, _ := block.DecodeFromJSON(json)
	// fmt.Println(b.Header.Height)
	// fmt.Println(b.Header.Size)
	// fmt.Println(b.Header.Timestamp)
	// fmt.Println(b.Header.Hash)
	// fmt.Println(b.Header.ParentHash)
	// fmt.Println(b.Value.Root())

	// json2, _ := b.EncodeToJSON()
	// fmt.Println(json2)
}
