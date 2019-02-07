package main

import (
	"fmt"

	"./p1"
)

func main() {
	//keyHexStr := hex.EncodeToString([]byte("j"))
	//fmt.Println(keyHexStr)
	//fmt.Println("hello world!")
	// fmt.Println("Encoded Result: %v", p1.Compact_encode(p1.StringToHexArray("ab")))
	//fmt.Println("Encoded Result: %v",p1.Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16}))
	fmt.Println("Encoded Result: %v", p1.Compact_encode([]uint8{2, 16}))
	//fmt.Println("Decoded Result: %v",p1.Compact_decode([]uint8{32}))
	//p1.Test_compact_encode();

	// mpt := p1.NewMPT()
	// mpt.Insert("a", "apple")
	// value, _, _ := mpt.GetKeyPath("a")
	// fmt.Println(string(value))
}
