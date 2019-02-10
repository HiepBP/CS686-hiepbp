package p1

import (
	"errors"
)

func (mpt *MerklePatriciaTrie) get_path(key string) (*Stack, error) {
	keyHexArr := string_to_hex_array(key)


	stack := new(Stack)

	// get the tree stored user data key to the value
	err := mpt.get_path_recursive(mpt.db[mpt.root], keyHexArr, 0, stack)

	return stack, err
}

func (mpt *MerklePatriciaTrie) get_path_recursive(root Node, key []uint8, pos int, stack *Stack) error {
	if root.is_empty() {
		return errors.New("problem: empty node")
	}
	stack.push(&root)
	switch root.node_type {
	case 0: //NULL
		return errors.New("problem: NULL node")
	case 1: //Branch
		if pos == len(key) { //everything matched
			//Get the value in branch node
			return nil
		}
		child_node_pointer := root.branch_value[key[pos]]
		if len(child_node_pointer) > 0 {
			child_node := mpt.db[child_node_pointer]
			if child_node.is_empty() {
				return errors.New("problem: can not find child")
			}
			return mpt.get_path_recursive(child_node, key, pos+1, stack)
		}
		return errors.New("problem: branch problem")
	case 2: //Ext or Leaf
		path := compact_decode(root.flag_value.encoded_prefix)
		if !root.is_leaf() { //Ext
			if len(path) > len(key)-pos || !path_compare(path, key[pos:pos+len(path)]) {
				return errors.New("problem: path not valid with this ext")
			}
			child_node_pointer := root.flag_value.value
			child_node := mpt.db[child_node_pointer]
			if child_node.is_empty() {
				return errors.New("problem: can not find child")
			}
			return mpt.get_path_recursive(child_node, key, pos+len(path), stack)
		} else { //Leaf, can not continue
			if len(path) != len(key)-pos || !path_compare(path, key[pos:pos+len(path)]) {
				return errors.New("problem: leaf path not match")
			}
			//Get the value in leaf node
			return nil
		}
	}
	return errors.New("problem: others")
}

func path_compare(arr1, arr2 []uint8) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
