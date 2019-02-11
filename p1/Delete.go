package p1

import (
	"errors"
)

func (mpt *MerklePatriciaTrie) Delete(key string) error {
	// TODO
	stack, err := mpt.get_path(key)
	if err == nil {
		return mpt.delete(key, stack)
	}
	return err
}

func (mpt *MerklePatriciaTrie) delete(key string, stack *Stack) error {
	key_hex_arr := string_to_hex_array(key)
	last_node := stack.pop()
	parent_node := stack.pop()
	var parent_node_hash = parent_node.hash_node()
	path := []uint8{}
	//Tree only has one leaf node
	if parent_node.is_empty() {
		delete(mpt.db, last_node.hash_node())
		mpt.root = ""
		return nil
	}

	switch last_node.node_type {
	case 0: //NULL
		return errors.New("problem: NULL node")
	case 1: //Branch
		delete(mpt.db, last_node.hash_node())
		last_node.branch_value[16] = ""
	case 2: //Leaf only, parent_node will be a branch
		path = compact_decode(last_node.flag_value.encoded_prefix)
		key_hex_arr = key_hex_arr[:len(key_hex_arr)-len(path)]
		delete(mpt.db, last_node.hash_node())
		branchIndex := key_hex_arr[len(key_hex_arr)-1]
		parent_node.branch_value[branchIndex] = ""
		//Update
		key_hex_arr = key_hex_arr[:len(key_hex_arr)-1]
		last_node = parent_node
		parent_node = stack.pop()
	}

	//Check if last_node(branch) only has 1 child or only contain value
	count := 0
	var new_path []uint8
	var new_value string
	branch_index := 0
	for i := 0; i < len(last_node.branch_value); i++ {
		if last_node.branch_value[i] != "" {
			branch_index = i
			count++
		}
	}
	if count > 1 && !parent_node.is_empty() { //Keep the branch
		stack.push(parent_node)
	} else { //Delete the branch and Update the currNode value to parent_node
		delete(mpt.db, parent_node_hash)
		var new_node Node
		switch parent_node.node_type {
		case 0: //NULL: last_node is the root
			if count == 1 {
				childNode := mpt.db[last_node.branch_value[branch_index]]
				delete(mpt.db, childNode.hash_node())
				new_path = compact_decode(childNode.flag_value.encoded_prefix)
				key_hex_arr = append(key_hex_arr, new_path...)
				new_path = append([]uint8{uint8(branch_index)}, new_path...)
				new_value = childNode.flag_value.value
				if childNode.is_leaf() {
					new_node = create_new_node(new_path, new_value, true)
				} else {
					new_node = create_new_node(new_path, new_value, false)
				}
				last_node = &new_node
			}

		case 1: //Branch(parent)
			parent_branch_index := key_hex_arr[len(key_hex_arr)-1]
			if branch_index == 16 {
				new_value = last_node.branch_value[16]
				new_node = create_new_node(new_path, new_value, true)
			} else { //Delete current child node from db then update the new path for that child
				child_node := mpt.db[last_node.branch_value[branch_index]]
				if child_node.node_type == 1 {
					new_path = append(new_path, uint8(branch_index))
					new_value = child_node.hash_node()
					new_node = create_new_node(new_path, new_value, false)
					key_hex_arr = append(key_hex_arr, uint8(branch_index))
				} else {
					delete(mpt.db, child_node.hash_node())
					new_path = compact_decode(child_node.flag_value.encoded_prefix)
					new_path = append([]uint8{uint8(branch_index)}, new_path...)
					key_hex_arr = append(key_hex_arr, new_path...)
					new_value = child_node.flag_value.value
					if child_node.is_leaf() {
						new_node = create_new_node(new_path, new_value, true)
					} else {
						new_node = create_new_node(new_path, new_value, false)
					}
				}
			}
			last_node = &new_node
			//Delete old parent_node before putting new one
			delete(mpt.db, parent_node.hash_node())
			parent_node.branch_value[parent_branch_index] = last_node.hash_node()
			stack.push(parent_node)
		case 2: //Ext(parent)
			ext_path := compact_decode(parent_node.flag_value.encoded_prefix)
			delete(mpt.db, parent_node.hash_node())
			if branch_index == 16 {
				new_path = ext_path
				new_value = last_node.branch_value[16]
				new_node = create_new_node(new_path, new_value, true)
			} else {
				child_node := mpt.db[last_node.branch_value[branch_index]]
				key_hex_arr = append(key_hex_arr, uint8(branch_index))
				new_path = append(ext_path, uint8(branch_index))
				if child_node.node_type == 1 {
					new_value = last_node.branch_value[branch_index]
					new_node = create_new_node(new_path, new_value, false)
				} else {
					if child_node.is_leaf() {
						delete(mpt.db, child_node.hash_node())
						key_hex_arr = append(key_hex_arr, compact_decode(child_node.flag_value.encoded_prefix)...)
						new_path = append(new_path, compact_decode(child_node.flag_value.encoded_prefix)...)
						new_value = child_node.flag_value.value
						new_node = create_new_node(new_path, new_value, true)
					} else {
						delete(mpt.db, child_node.hash_node())
						key_hex_arr = append(key_hex_arr, compact_decode(child_node.flag_value.encoded_prefix)...)
						new_path = append(new_path, compact_decode(child_node.flag_value.encoded_prefix)...)
						new_value = child_node.flag_value.value
						new_node = create_new_node(new_path, new_value, false)
					}
				}
			}
			last_node = &new_node
		}
	}
	stack.push(last_node)
	mpt.update_hash_for_delete(key_hex_arr, stack)
	return nil
}
