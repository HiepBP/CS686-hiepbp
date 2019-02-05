package p1

//func (mpt *MerklePatriciaTrie) getWithNode(key string) (string, *Stack, error) {
//	keyHexStr := hex.EncodeToString([]byte(key))
//
//	fmt.Println("Merkle get %s ...", keyHexStr)
//
//	s := new(Stack)
//
//	// get the tree stored user data key to the value
//	userValue, err := mt.findValue(mt.root, keyHexStr, 0, s)
//	if err != nil {
//		log.Error("Error getting user data from m", err)
//		return nil, s, err
//	}
//
//	if userValue == nil {
//		log.Debug("No data in Merkle tree for %s", keyHexStr)
//		return nil, s, nil
//	}
//
//	log.Debug("Found %s value in Merkle tree for key: %s", hex.EncodeToString(userValue), keyHexStr)
//
//	// pull the data from the user data store
//	value, err := mt.userData.Get(userValue, nil)
//
//	if err == leveldb.ErrNotFound {
//		// the value from the merkle tree is the short user value - return it
//		return userValue, s, nil
//	}
//
//	if err != nil {
//		return nil, s, err
//	}
//
//	// return actual user value
//	return value, s, err
//}
