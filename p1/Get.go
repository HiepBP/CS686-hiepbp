package p1

import (
	"errors"
)

func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	log.Debug("Get: ", key)
	if len(mpt.root) == 0 {
		return "", errors.New("empty trie")
	}
	key_hex_array := string_to_hex_array(key)
	return mpt.get(mpt.db[mpt.root], key_hex_array, 0)
}

func (mpt *MerklePatriciaTrie) get(root Node, key []uint8, pos int) (string, error) {
	if root.is_empty() {
		return "", errors.New("problem: empty node")
	}

	switch root.node_type {
	case 0: //NULL
		return "", errors.New("problem: NULL node")
	case 1: //Branch
		if pos == len(key) { //everything matched
			//Get the value in branch node
			return root.branch_value[16], nil
		}
		child_node_pointer := root.branch_value[key[pos]]
		if len(child_node_pointer) > 0 {
			child_node := mpt.db[child_node_pointer]
			if child_node.is_empty() {
				return "", errors.New("problem: can not find child")
			}
			return mpt.get(child_node, key, pos+1)
		}
		return "", errors.New("problem: branch problem")
	case 2: //Ext or Leaf
		path := compact_decode(root.flag_value.encoded_prefix)
		if !root.is_leaf() { //Ext
			if len(path) > len(key)-pos || !path_compare(path, key[pos:pos+len(path)]) {
				return "", errors.New("problem: ext path not match")
			}
			child_node_pointer := root.flag_value.value
			child_node := mpt.db[child_node_pointer]
			if child_node.is_empty() {
				return "", errors.New("problem: can not find child")
			}
			return mpt.get(child_node, key, pos+len(path))
		} else { //Leaf
			if len(path) != len(key)-pos || !path_compare(path, key[pos:pos+len(path)]) {
				return "", errors.New("problem: leaf path not match")
			}
			//Get the value in leaf node
			return root.flag_value.value, nil
		}
	}
	return "", errors.New("problem: others")
}
