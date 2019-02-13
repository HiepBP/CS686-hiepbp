package p1

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	log.Debug("Insert ", key, " : ", new_value)
	// hexKey := hex.EncodeToString([]byte(key))
	// fmt.Println("Insert data with key: ", hexKey)
	stack, _ := mpt.get_path(key)
	length := get_path_length(string_to_hex_array(key), stack)
	mpt.insert(length, key, new_value, stack)
}

//pos: number of match path
//stack: path from root to inserted place
func (mpt *MerklePatriciaTrie) insert(pos int, key string, value string, stack *Stack) {
	key_hex_arr := string_to_hex_array(key)
	//Empty tree
	if stack.length == 0 {
		new_node := create_new_node(key_hex_arr, value, true)
		new_node_hash := new_node.hash_node()
		stack.push(&new_node)
		mpt.db[new_node_hash] = new_node
		mpt.root = new_node.hash_node()
		return
	}
	//Check the last node in the path stack to handle the insertion
	last_node := stack.pop()
	last_node_hash := last_node.hash_node()
	switch last_node.node_type {

	case 1: //LastNode is a branch
		stack.push(last_node)
		if pos-1 == len(key_hex_arr) { //key match the whole path to branch, update the value of branch
			last_node.branch_value[16] = value
			//remove it from db
			delete(mpt.db, last_node_hash)
		} else { //add new leaf(key[pos:], value) to trie and db
			// pos++
			new_node := create_new_node(key_hex_arr[pos:], value, true)
			stack.push(&new_node)
		}
		mpt.update_node_hash_value(key_hex_arr, stack)
		return

	case 2: //Last node is a leaf/extension
		//Get the common prefix part in the path
		if last_node.is_leaf() {
			last_node_path := compact_decode(last_node.flag_value.encoded_prefix)
			last_node_path_length := get_path_length(key_hex_arr, stack)
			common_prefix := get_common_prefix(last_node_path, key_hex_arr[last_node_path_length:])
			//If it is leaf and its path same with the new path
			if len(common_prefix) == len(last_node_path) && //Common prefix equal with last_node_path
				pos+len(common_prefix) == len(key_hex_arr) { //Pos + common_prefix == keyPath
				delete(mpt.db, last_node_hash)
				last_node.flag_value.value = value
				stack.push(last_node)
				mpt.update_node_hash_value(key_hex_arr, stack)
				return
			}
		} else {
			pos -= len(compact_decode(last_node.flag_value.encoded_prefix))
		}

		// If have common, create ext -> branch
		// If no common, create branch
		last_node_path := compact_decode(last_node.flag_value.encoded_prefix)
		common_prefix := get_common_prefix(last_node_path, key_hex_arr[pos:])
		if len(common_prefix) > 0 {
			ext_path := last_node_path[:len(common_prefix)]
			new_ext_node := create_new_node(ext_path, "", false)
			stack.push(&new_ext_node)
			if len(common_prefix) < len(last_node_path) {
				last_node_path = last_node_path[len(common_prefix):]
			} else {
				last_node_path = nil
			}
			//Increase the pos (number of match path) base on common prefix
			pos += len(common_prefix)
		}

		//Then add a new branch
		new_branch_node := Node{node_type: 1}
		stack.push(&new_branch_node)

		//Check if last node in the stack is already covered in the ext or not
		if len(last_node_path) > 0 {
			//Get the index in branch
			branch_index := last_node_path[0]
			last_node_path = last_node_path[1:]
			//Create new extension or leaf from last_node in path
			if len(last_node_path) > 0 || last_node.is_leaf() {
				//Delete old node from db, update the encoded_prefix, add to branch, than add new one to db
				delete(mpt.db, last_node_hash)
				if last_node.is_leaf() {
					last_node.flag_value.encoded_prefix = compact_encode(append(last_node_path, 16))
				} else {
					last_node.flag_value.encoded_prefix = compact_encode(last_node_path)
				}
				new_branch_node.branch_value[branch_index] = last_node.hash_node()
				mpt.db[last_node.hash_node()] = *last_node
			} else { //Last node is ext and len = 0
				delete(mpt.db, last_node_hash)
				new_branch_node.branch_value[branch_index] = last_node.flag_value.value
			}
		} else { //If the length is 0, remove the ext/leaf and add new leaf to the branch
			delete(mpt.db, last_node_hash)
			if last_node.is_leaf() {
				new_branch_node.branch_value[16] = last_node.flag_value.value
			}
		}
		if pos < len(key_hex_arr) {
			pos++
			if pos > len(key_hex_arr) {
				pos = len(key_hex_arr)
			}
			new_leaf_node := create_new_node(key_hex_arr[pos:], value, true)
			stack.push(&new_leaf_node)
		} else {
			new_branch_node.branch_value[16] = value
		}
		mpt.update_node_hash_value(key_hex_arr, stack)
		return
	}
	return
}
