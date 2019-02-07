package p1

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
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

func NewMPT() *MerklePatriciaTrie {
	db := make(map[string]Node)
	root := ""
	return &MerklePatriciaTrie{db, root}
}

func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	// TODO
	return "", errors.New("path_not_found")
}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
	hexKey := hex.EncodeToString([]byte(key))
	fmt.Println("Insert data with key: ", hexKey)
	_, stack, _ := mpt.GetKeyPath(key)
	length := GetPathLength(stack)
	mpt.InsertAtTheEndOfPath(length, key, new_value, stack)

}

//pos: number of match path
//stack: path from root to inserted place
func (mpt *MerklePatriciaTrie) InsertAtTheEndOfPath(pos int, key string, value string, stack *Stack) {
	//Empty tree
	if stack.length == 0 {
		newNode := CreateNewNodeWithString(key, value, true)
		hashedNode := newNode.hash_node()
		mpt.db[hashedNode] = newNode
		mpt.root = hashedNode
		return
	}
	//Check the last node in the path stack to handle the insertion
	lastNode := stack.pop()
	lastNodeHashed := lastNode.hash_node()
	switch lastNode.node_type {

	case 1: //LastNode is a branch
		stack.push(lastNode)
		if pos == len(key) { //key match the whole path to branch, update the value of branch
			lastNode.branch_value[16] = value
			//remove it from db
			delete(mpt.db, lastNodeHashed)
		} else { //add new leaf(key[pos:], value) to trie and db
			pos++
			newNode := CreateNewNodeWithString(key[pos:], value, true)
			stack.push(newNode)
		}
		mpt.UpdateNodeAndHashValue(key, stack)
		return

	case 2: //Last node is a leaf/extension
		//Get the common prefix part in the path
		keyHexArr := StringToHexArray(key)
		lastNodePath := Compact_decode(lastNode.flag_value.encoded_prefix)
		lastNodePathLength := GetPathLength(stack)
		commonPrefix := GetCommonPrefixPath(lastNodePath, keyHexArr[lastNodePathLength:])

		//If it is leaf and its path same with the new path
		if isLeaf(lastNode.flag_value.encoded_prefix) &&
			len(commonPrefix) == len(lastNodePath) &&
			pos == len(keyHexArr) {
			delete(mpt.db, lastNodeHashed)
			lastNode.flag_value.value = value
			stack.push(lastNode)
			mpt.UpdateNodeAndHashValue(key, stack)
			return
		}

		// Leaf/Extension with different in path
		// If have common, create ext -> branch
		// If no common, create branch
		// Check if it has common, create ext or not
		if len(commonPrefix) > 0 {
			extKey := lastNodePath[:len(commonPrefix)]
			newExtNode := CreateNewNodeWithHexArray(extKey, "", false)
			stack.push(newExtNode)
			if len(commonPrefix) < lastNodePathLength {
				lastNodePath = lastNodePath[len(commonPrefix):]
			} else {
				lastNodePath = nil
			}
			//Increase the pos (number of match path) base on common prefix
			pos += len(commonPrefix)
		}

		//Then add a new branch
		newBranchNode := Node{node_type: 1}
		stack.push(newBranchNode)

		//Check if last node in the stack is already covered in the ext or not
		if len(lastNodePath) > 0 {
			//Get the index in branch
			brandIndex := lastNodePath[0]
			lastNodePath = lastNodePath[1:]
			//Create new extension or leaf from lastNode in path
			if len(lastNodePath) > 0 || isLeaf(lastNodePath) {
				//Delete old node from db, update the encoded_prefix, add to branch, than add new one to db
				delete(mpt.db, lastNodeHashed)
				lastNode.flag_value.encoded_prefix = Compact_encode(lastNodePath)
				newBranchNode.branch_value[brandIndex] = lastNode.hash_node()
				mpt.db[lastNode.hash_node()] = lastNode
			} else { //Last node is ext and len = 0
				delete(mpt.db, lastNodeHashed)
				newBranchNode.branch_value[brandIndex] = lastNode.flag_value.value
			}
		} else { //If the length is 0, remove the ext/leaf and add new leaf to the branch
			delete(mpt.db, lastNodeHashed)
			if isLeaf(lastNode.flag_value.encoded_prefix) {
				newBranchNode.branch_value[16] = lastNode.flag_value.value
			}
		}
		if pos < len(key) {
			pos++
			newLeafNode := CreateNewNodeWithString(key[pos:], value, true)
			stack.push(newLeafNode)
		}
		mpt.UpdateNodeAndHashValue(key, stack)
		return
	}
	return
}

func (mpt *MerklePatriciaTrie) Delete(key string) error {
	// TODO
	return errors.New("path_not_found")
}

func Compact_encode(hex_array []uint8) []uint8 {
	var term = 0
	var result []uint8
	//Check the last value in array
	if hex_array[len(hex_array)-1] == 16 {
		term = 1
	}
	if term == 1 {
		hex_array = hex_array[:len(hex_array)-1]
	}
	var oddlen = len(hex_array) % 2
	var flags uint8 = uint8(2*term + oddlen)
	if oddlen == 1 {
		hex_array = append([]uint8{flags}, hex_array...)
	} else {
		hex_array = append([]uint8{flags, 0}, hex_array...)
	}
	for i := 0; i < len(hex_array); i += 2 {
		result = append(result, 16*hex_array[i]+hex_array[i+1])
	}
	return result
}

// If Leaf, ignore 16 at the end
func Compact_decode(encoded_arr []uint8) []uint8 {
	var result []uint8
	//Decode back the encoded_path
	for i := 0; i < len(encoded_arr); i++ {
		result = append(result, encoded_arr[i]/16)
		result = append(result, encoded_arr[i]%16)
	}
	//Check if it is even or odd len
	if result[0] == 1 || result[0] == 3 {
		result = result[1:len(result)]
	} else {
		result = result[2:len(result)]
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
