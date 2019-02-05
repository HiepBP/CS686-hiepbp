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
	value          string
}

type Node struct {
	node_type    int // 0: Null, 1: Branch, 2: Ext or Leaf
	branch_value [17]string
	flag_value   Flag_value
}

func (node Node) isEmpty() bool {
	return reflect.DeepEqual(node, Node{})
}

type MerklePatriciaTrie struct {
	db   map[string]Node
	root string
}

func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	// TODO
	return "", errors.New("path_not_found")
}

//func (mpt *MerklePatriciaTrie) get_with_stack(key string) (string, Stack, error){
//	hexKey := hex.EncodeToString([]byte(key))
//	fmt.Println("Get value for: %s", hexKey)
//}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
	hexKey := hex.EncodeToString([]byte(key))
	fmt.Println("Insert data with key: %s", hexKey)
}

func (mpt *MerklePatriciaTrie) Delete(key string) error {
	// TODO
	return errors.New("path_not_found")
}

func insert_node(key string, new_value string, db map[string]Node, currNode Node) {
	switch currNode.node_type {
	case 0: //Null
		break
	case 1: //Branch
		if len(key) == 0 {
			currNode.branch_value[16] = new_value
		} else {
			var firstIndex = key[0]
			if len(currNode.branch_value[firstIndex]) == 0 {
				//TODO: Insert new leaf and add hashcode
			} else {
				var child_node = db[currNode.branch_value[firstIndex]]
				insert_node(key[1:], new_value, db, child_node)
				currNode.hash_node()
			}
		}
	case 2: //Ext or Leaf
		break
	}
}

func Compact_encode(hex_array []uint8) []uint8 {
	var term = 0
	var result []uint8
	if hex_array[len(hex_array)-1] == 16 {
		term = 1
	}
	if term == 1 {
		hex_array = hex_array[0 : len(hex_array)-1]
	}
	var oddlen = len(hex_array) % 2
	var flags uint8 = uint8(2*term + oddlen)
	if oddlen == 1 {
		hex_array = append([]uint8{flags}, hex_array...)
	} else {
		hex_array = append([]uint8{flags, 0}, hex_array...)
	}
	fmt.Println("Hex Array: %v", hex_array)
	for i := 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	return result
}

// If Leaf, ignore 16 at the end
func compact_decode(encoded_arr []uint8) []uint8 {
	var result []uint8
	for i := 0; i < len(encoded_arr); i++ {
		result = append(result, encoded_arr[i]/16)
		result = append(result, encoded_arr[i]%16)
	}
	fmt.Println("Decoded array: %v", result)
	if result[0] == 0 {
		result = result[2:len(result)]
	} else if result[0] == 1 {
		result = result[1:len(result)]
	} else if result[0] == 2 {
		if len(result) == 2 {
			result = result[:0]
		} else {
			result = result[2:len(result)]
		}
	} else if result[0] == 3 {
		result = result[1:len(result)]
	}
	return result
}

//func Test_compact_encode() {
//	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
//	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
//	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
//	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
//}

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

func StringToHexArray(str string) []uint8 {
	var hexArrayResult []uint8
	for i := 0; i < len(str); i++ {
		hexArrayResult = append(hexArrayResult, str[i]/16)
		hexArrayResult = append(hexArrayResult, str[i]%16)
	}
	return hexArrayResult
}
