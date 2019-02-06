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
