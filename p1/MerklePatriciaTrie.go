package p1

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

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

func (node Node) is_empty() bool {
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

func (mpt *MerklePatriciaTrie) Initial() {
	mpt.db = make(map[string]Node)
	mpt.root = ""
}

func compact_encode(hex_array []uint8) []uint8 {
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
func compact_decode(encoded_arr []uint8) []uint8 {
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

func test_compact_encode() {
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
}

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

func (mpt *MerklePatriciaTrie) Get_db_length() int {
	return len(mpt.db)
}

//Support function

func (node *Node) String() string {
	str := "empty string"
	switch node.node_type {
	case 0:
		str = "[Null Node]"
	case 1:
		str = "Branch["
		for i, v := range node.branch_value[:16] {
			str += fmt.Sprintf("%d=\"%s\", ", i, v)
		}
		str += fmt.Sprintf("value=%s]", node.branch_value[16])
	case 2:
		encoded_prefix := node.flag_value.encoded_prefix
		node_name := "Leaf"
		if is_ext_node(encoded_prefix) {
			node_name = "Ext"
		}
		ori_prefix := strings.Replace(fmt.Sprint(compact_decode(encoded_prefix)), " ", ", ", -1)
		str = fmt.Sprintf("%s<%v, value=\"%s\">", node_name, ori_prefix, node.flag_value.value)
	}
	return str
}

func is_ext_node(encoded_arr []uint8) bool {
	return encoded_arr[0]/16 < 2
}

func node_to_string(node Node) string {
	return node.String()
}

func (mpt *MerklePatriciaTrie) String() string {
	content := fmt.Sprintf("ROOT=%s\n", mpt.root)
	for hash := range mpt.db {
		content += fmt.Sprintf("%s: %s\n", hash, node_to_string(mpt.db[hash]))
	}
	return content
}

func (mpt *MerklePatriciaTrie) Order_nodes() string {
	raw_content := mpt.String()
	content := strings.Split(raw_content, "\n")
	root_hash := strings.Split(strings.Split(content[0], "HashStart")[1], "HashEnd")[0]
	queue := []string{root_hash}
	i := -1
	rs := ""
	cur_hash := ""
	for len(queue) != 0 {
		last_index := len(queue) - 1
		cur_hash, queue = queue[last_index], queue[:last_index]
		i += 1
		line := ""
		for _, each := range content {
			if strings.HasPrefix(each, "HashStart"+cur_hash+"HashEnd") {
				line = strings.Split(each, "HashEnd: ")[1]
				rs += each + "\n"
				rs = strings.Replace(rs, "HashStart"+cur_hash+"HashEnd", fmt.Sprintf("Hash%v", i), -1)
			}
		}
		temp2 := strings.Split(line, "HashStart")
		flag := true
		for _, each := range temp2 {
			if flag {
				flag = false
				continue
			}
			queue = append(queue, strings.Split(each, "HashEnd")[0])
		}
	}
	return rs
}

func TestCompact() {
	test_compact_encode()
}

func (mpt *MerklePatriciaTrie) update_node_hash_value(key []uint8, stack *Stack) {
	var pre_node Node
	var cur_node Node
	var pos = len(key) - 1
	items := stack.retrieve()
	// for item := stack.top; item != nil; item = item.next {
	for i := 0; i < len(items); i++ {
		cur_node = items[i]
		delete(mpt.db, cur_node.hash_node())
		switch cur_node.node_type {
		case 1: //Branch
			if !pre_node.is_empty() {
				cur_node.branch_value[key[pos]] = pre_node.hash_node()
				pos--
			}
		case 2: //Ext/leaf
			pos -= len(compact_decode(cur_node.flag_value.encoded_prefix))
			if cur_node.is_leaf() {

			} else {
				if !pre_node.is_empty() {
					cur_node.flag_value.value = pre_node.hash_node()
				}
			}
		}
		mpt.db[cur_node.hash_node()] = cur_node
		pre_node = cur_node
	}
	mpt.root = cur_node.hash_node()
	return
}
