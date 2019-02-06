package p1

func isLeaf(encodedPrefix []uint8) bool {
	if encodedPrefix[0]/16 == 0 || encodedPrefix[0]/16 == 1 {
		return false
	}
	return true
}

func StringToHexArray(str string) []uint8 {
	var hexArrayResult []uint8
	for i := 0; i < len(str); i++ {
		hexArrayResult = append(hexArrayResult, str[i]/16)
		hexArrayResult = append(hexArrayResult, str[i]%16)
	}
	return hexArrayResult
}

func (mpt *MerklePatriciaTrie) GetPathLength(s *Stack) int {
	length := 0
	nodePath := s.retrieve()
	for _, node := range nodePath {
		if node.node_type == 1 {
			length++
		} else if !isLeaf(node.flag_value.encoded_prefix) { //If it is EXT, increase the length equal with the decoded_prefix
			length += len(Compact_decode(node.flag_value.encoded_prefix))
		}
	}
	return length
}