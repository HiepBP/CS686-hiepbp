package p1

import "errors"

//Check if the Node if Leaf/Extension base on encodedPrefix
func (node *Node) is_leaf() bool {
	if node.flag_value.encoded_prefix[0]/16 == 0 ||
		node.flag_value.encoded_prefix[0]/16 == 1 {
		return false
	}
	return true
}

//Convert from string to hexadecimal array
func string_to_hex_array(str string) []uint8 {
	var hexArrayResult []uint8
	for i := 0; i < len(str); i++ {
		hexArrayResult = append(hexArrayResult, str[i]/16)
		hexArrayResult = append(hexArrayResult, str[i]%16)
	}
	return hexArrayResult
}

//Get the length of inserted node path on the trie
func get_path_length(key []uint8, s *Stack) int {
	length := 0
	nodePath := s.retrieve()
	for i := len(nodePath) - 1; i >= 0; i-- {
		node := nodePath[i]
		if node.node_type == 1 {
			length++
		} else if !node.is_leaf() { //If it is EXT, increase the length equal with the decoded_prefix
			length += len(compact_decode(node.flag_value.encoded_prefix))
			// length += len(get_common_prefix(key[length:], compact_decode(node.flag_value.encoded_prefix)))
		}
	}
	return length
}

func create_new_node(key []uint8, value string, isLeaf bool) Node {
	if isLeaf {
		key = append(key, 16) //append 16 to leaf
	}
	flagValue := Flag_value{
		encoded_prefix: compact_encode(key),
		value:          value,
	}
	return Node{
		node_type:  2,
		flag_value: flagValue,
	}
}

func get_common_prefix(path1, path2 []uint8) []uint8 {
	var result []uint8
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] == path2[i] {
			result = append(result, path1[i])
		} else {
			break
		}
	}
	return result
}

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
