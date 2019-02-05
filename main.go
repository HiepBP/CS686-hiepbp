package main

import (
	"encoding/hex"
	"fmt"
	"./p1"
)

func main() {
	keyHexStr := hex.EncodeToString([]byte("j"))
	fmt.Println(keyHexStr)
	fmt.Println("hello world!")
	//fmt.Println("Encoded Result: %v",p1.Compact_encode([]uint8{1, 2, 3, 4, 5}))
	//fmt.Println("Encoded Result: %v",p1.Compact_encode([]uint8{0 ,1, 2, 3, 4, 5}))
	//fmt.Println("Encoded Result: %v",p1.Compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16}))
	//fmt.Println("Encoded Result: %v",p1.Compact_encode([]uint8{15, 1, 12, 11, 8, 16}))
	//fmt.Println("Decoded Result: %v",p1.Compact_decode([]uint8{63, 28, 184}))
	p1.Test_compact_encode();
}
