package main

import (
	"fmt"

	"./p2"
)

func main() {

	jsonBlockChain := "[{\"height\":1,\"timeStamp\":1551025401,\"hash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"parentHash\":\"genesis\",\"size\":2089,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"f367b7f59c651e69be7e756298aad62fb82fddbfeda26cb06bfd8adf9c8aa094\",\"parentHash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"05ac44dd82b6cc398a5e9664add21856ae19d107d9035af5fc54c9b0ffdef336\",\"parentHash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}}]"
	bc, err := p2.DecodeJsonToBlockChain(jsonBlockChain)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bc.EncodeToJSON()
	if err != nil {
		fmt.Println(err)
	}

	// mpt := p1.MerklePatriciaTrie{}
	// mpt.Initial()
	// mpt.Insert("hello", "world")
	// mpt.Insert("charles", "ge")
	// b1 := p2.Initial(1, 1234567890, "genesis", mpt)
	// b2 := p2.Initial(2, 1234567890, b1.Header.Hash, mpt)
	// bc := p2.NewBlockChain()
	// bc.Insert(b1)
	// bc.Insert(b2)
	// jsonB1, _ := b1.EncodeToJSON()
	// jsonB2, _ := b2.EncodeToJSON()
	// fmt.Println(jsonB1)
	// fmt.Println(jsonB2)
	// fmt.Println(bc.Length())

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
