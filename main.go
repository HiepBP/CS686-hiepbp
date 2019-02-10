package main

import (
	"fmt"

	"./p1"
)

func main() {
	//keyHexStr := hex.EncodeToString([]byte("j"))
	//fmt.Println(keyHexStr)
	//fmt.Println("hello world!")
	// fmt.Println("Encoded Result: %v", p1.compact_encode(p1.string_to_hex_array("ab")))
	//fmt.Println("Encoded Result: %v",p1.compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16}))
	// fmt.Println("Encoded Result: %v", p1.compact_encode([]uint8{2, 16}))
	//fmt.Println("Decoded Result: %v",p1.compact_decode([]uint8{32}))
	//p1.Test_compact_encode();
	mpt := p1.NewMPT()

	//1
	// mpt.Insert("a", "apple")
	// mpt.Insert("b", "banana")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Delete("a")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("a"))
	// fmt.Println(mpt.Get("b"))

	//2
	// mpt.Insert("a", "apple")
	// mpt.Insert("p", "banana")
	// mpt.Insert("abc", "new")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("p"))

	//3
	// mpt.Insert("p", "apple")
	// mpt.Insert("aaaaa", "banana")
	// mpt.Insert("aaaap", "orange")
	// mpt.Insert("aa", "new")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Delete("aaaap"))
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("aa"))

	//4
	// mpt.Insert("p", "apple")
	// mpt.Insert("aa", "banana")
	// mpt.Insert("ap", "orange")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Delete("p"))
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("ap"))
}
