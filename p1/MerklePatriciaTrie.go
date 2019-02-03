package p1

import (
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	"reflect"
)

type Flag_value struct {
	encoded_prefix []uint8
	value string
}

type Node struct {
	node_type int // 0: Null, 1: Branch, 2: Ext or Leaf
	branch_value [17]string
	flag_value Flag_value
}

type MerklePatriciaTrie struct {
	db map[string]Node
	root string
}

func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	// TODO
	return "", errors.New("path_not_found")
}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
}

func (mpt *MerklePatriciaTrie) Delete(key string) error {
	// TODO
	return errors.New("path_not_found")
}

func compact_encode(hex_array []uint8) []uint8 {
	var term = 0;
	var result []uint8;
	if(hex_array[len(hex_array)-1] == 16){
		term = 1;
	}
	if term == 1{
		hex_array = hex_array[0:len(hex_array)-1];
	}
	var oddlen = len(hex_array) % 2;
	var flags uint8 = uint8(2 * term + oddlen);
	if oddlen == 1 {
		hex_array = append([]uint8{flags}, hex_array...);
	} else{
		hex_array = append([]uint8{flags, 0}, hex_array...);
	}
	fmt.Println("Hex Array: %v", hex_array);
	for i:=0; i<len(hex_array); i+=2{
		result = append(result, 16*hex_array[i] + hex_array[i+1]);
	}
	return result
}

// If Leaf, ignore 16 at the end
func compact_decode(encoded_arr []uint8) []uint8 {
	var result []uint8;
	for i:=0; i<len(encoded_arr); i++ {
		result = append(result, encoded_arr[i] / 16);
		result = append(result, encoded_arr[i] % 16);
	}
	println("{:?} ",result);
	if result[0] == 0 {
		result = result[2:len(result)];
	} else  if result[0] == 1 {
		result = result[1:len(result)];
	} else  if result[0] == 2 {
		if len(result) == 2{
			result = result[:0];
		} else {
			result = result[2:len(result)-1];
		}
	} else if result[0] == 3 {
		result = result[1:len(result)-1];
	}
	return result
}

func test_compact_encode() {
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
}

func (node *Node) hash_node() string {
	var str string
	switch node.node_type {
	case 0:
		str = ""
	case 1:
		str = "branch_"
		for _, v := range node.branch_value {
			str += v
		}
	case 2:
		str = node.flag_value.value
	}

	sum := sha3.Sum256([]byte(str))
	return "HashStart_" + hex.EncodeToString(sum[:]) + "_HashEnd"
}