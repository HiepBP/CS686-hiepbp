package p1

import (
	"errors"
	"fmt"
)

func (mpt *MerklePatriciaTrie) GetKeyPath(key string) (*Stack, error) {
	keyHexArr := StringToHexArray(key)

	fmt.Println("Merkle get ...", keyHexArr)

	stack := new(Stack)

	// get the tree stored user data key to the value
	err := mpt.getThroughPath(mpt.db[mpt.root], keyHexArr, 0, stack)
	if err != nil {
		fmt.Println(err)
		return stack, err
	}

	return stack, err
}

func (mpt *MerklePatriciaTrie) getThroughPath(root Node, k []uint8, pos int, stack *Stack) error {
	if root.isEmpty() {
		return nil
	}
	stack.push(root)
	switch root.node_type {
	case 0: //NULL
		return nil
	case 1: //Branch
		if pos == len(k) { //everything matched
			//Get the value in branch node
			return nil
		}
		childNodePointer := root.branch_value[pos]
		if len(childNodePointer) > 0 {
			childNode := mpt.db[childNodePointer]
			if childNode.isEmpty() {
				return errors.New("problem: can not find child")
			}
			return mpt.getThroughPath(childNode, k, pos+1, stack)
		}
		return nil
	case 2: //Ext or Leaf
		path := Compact_decode(root.flag_value.encoded_prefix)
		if !root.isLeaf() { //Ext
			if len(path) > len(k)-pos || !pathCompare(path, k[pos:pos+len(path)]) {
				return nil
			}
			childNodePointer := root.flag_value.value
			childNode := mpt.db[childNodePointer]
			if childNode.isEmpty() {
				return errors.New("problem: can not find child")
			}
			return mpt.getThroughPath(childNode, k, pos+len(path), stack)
		} else { //Leaf
			if len(path) > len(k)-pos || !pathCompare(path, k[pos:pos+len(path)]) {
				return nil
			}
			//Get the value in leaf node
			return nil
		}
	}
	return nil
}

func pathCompare(arr1, arr2 []uint8) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
