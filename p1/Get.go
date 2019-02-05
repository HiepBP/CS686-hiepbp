package p1

import (
	"errors"
	"fmt"
)

func (mpt *MerklePatriciaTrie) getWithNode(key string) (string, *Stack, error) {
	keyHexArr := StringToHexArray(key)

	fmt.Println("Merkle get %s ...", keyHexArr)

	s := new(Stack)
	
	// get the tree stored user data key to the value
	value, err := mpt.getValue(mpt.db[mpt.root], keyHexArr, 0, s)
	if err != nil {
		fmt.Println(err)
		return "", s, err
	}

	if value == "" {
		fmt.Println("No data in Merkle tree for %s", keyHexArr)
		return "", s, nil
	}

	fmt.Println("Found value %s in Merkle tree for key: %s", value, keyHexArr)

	// return alue
	return value, s, err
}

func (mpt *MerklePatriciaTrie) getValue(root Node, k []uint8, pos int, s *Stack) (string, error) {

	s.push(root)
	switch root.node_type {
	case 0: //NULL
		return "", nil
	case 1: //Branch
		if pos == len(k) { //everything matched
			//Get the value in branch node
			return root.branch_value[16], nil
		}
		childNodePointer := root.branch_value[pos]
		if len(childNodePointer) > 0 {
			childNode := mpt.db[childNodePointer]
			if childNode.isEmpty() {
				return "", errors.New("problem: can not find child")
			}
			return mpt.getValue(childNode, k, pos+1, s)
		}
		return "", nil
	case 2: //Ext or Leaf
		path := compact_decode(root.flag_value.encoded_prefix)
		if !isLeaf(root.flag_value.encoded_prefix) { //Ext
			if len(path) > len(k)-pos || !pathCompare(path, k[pos:pos+len(path)]) {
				return "", nil
			}
			childNodePointer := root.flag_value.value
			childNode := mpt.db[childNodePointer]
			if childNode.isEmpty() {
				return "", errors.New("problem: can not find child")
			}
			return mpt.getValue(childNode, k, pos+len(path), s)
		} else { //Leaf
			if len(path) > len(k)-pos || !pathCompare(path, k[pos:pos+len(path)]) {
				return "", nil
			}
			//Get the value in leaf node
			return root.flag_value.value, nil
		}
	}
	return "", nil

}

func pathCompare(arr1, arr2 []uint8) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func isLeaf(encodedPrefix []uint8) bool {
	if encodedPrefix[0]/16 == 0 || encodedPrefix[0]/16 == 1 {
		return false
	}
	return true
}
