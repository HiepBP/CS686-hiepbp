package main

import (
	"fmt"

	"./p1"
)

func main() {
	mpt := p1.NewMPT()

	//1
	// mpt.Insert("p", "apple")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aa", "banana")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("ap", "orange")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("b", "new")
	// fmt.Println(mpt.Get("b"))
	// mpt.Delete("b")
	// fmt.Println(mpt.Get("p"))
	// fmt.Println(mpt.Get("aa"))
	// fmt.Println(mpt.Get("ap"))
	// fmt.Println(mpt.Get("b"))
	// fmt.Println(mpt.Get_db_length())

	//2
	// mpt.Insert("aaa", "apple")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aap", "banana")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("bb", "right leaf")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aa", "new")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Delete("aa")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("aaa"))
	// fmt.Println(mpt.Get("aap"))
	// fmt.Println(mpt.Get("bb"))
	// fmt.Println(mpt.Get("aa"))

	//031
	// mpt.Insert("p", "apple")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aaa", "banana")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aap", "orange")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("b", "new")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Delete("b")
	// fmt.Println(mpt.Get_db_length())

	//111
	// mpt.Insert("aa", "apple")
	// mpt.Insert("ap", "banana")
	// mpt.Insert("b", "new")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Delete("b")
	// fmt.Println(mpt.Get_db_length())

	//140
	// mpt.Insert("p", "apple")
	// mpt.Insert("aaaa", "banana")
	// mpt.Insert("aaap", "orange")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("a", "new")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Delete("a")
	// fmt.Println(mpt.Get_db_length())
	// fmt.Println(mpt.Get("p"))
	// fmt.Println(mpt.Get("aaaa"))
	// fmt.Println(mpt.Get("aaap"))
	// fmt.Println(mpt.Get("a"))

	//030

	// mpt.Insert("aaa", "apple")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aap", "banana")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("bb", "right leaf")
	// fmt.Println(mpt.Get_db_length())
	// mpt.Insert("aa", "new")
	// fmt.Println(mpt.Get_db_length())

	//v_np

	mpt.Insert("a", "old")
	fmt.Println(mpt.Get_db_length())
	mpt.Insert("aa", "apple")
	fmt.Println(mpt.Get_db_length())
	mpt.Insert("ap", "banana")
	fmt.Println(mpt.Get_db_length())
	mpt.Insert("a", "new")
	fmt.Println(mpt.Get_db_length())
	mpt.Delete("a")
	fmt.Println(mpt.Get_db_length())
}
