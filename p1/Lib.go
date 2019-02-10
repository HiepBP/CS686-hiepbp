package p1

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
func get_path_length(s *Stack) int {
	length := 0
	nodePath := s.retrieve()
	for _, node := range nodePath {
		if node.node_type == 1 {
			length++
		} else if !node.is_leaf() { //If it is EXT, increase the length equal with the decoded_prefix
			length += len(compact_decode(node.flag_value.encoded_prefix))
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
