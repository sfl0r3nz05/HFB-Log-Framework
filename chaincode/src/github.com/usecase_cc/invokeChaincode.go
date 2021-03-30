//	package main
//	
//	import (
//		"fmt"
//		"crypto/sha256"
//		"encoding/base64"
//	)

//	func invokeChaincode(uuid string, TxID string, CHANNEL_ENV string) {
//		hasher := sha256.New()
//		hasher.Write([]byte(TxID))
//		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
//		params := []string{"set", uuid, sha}
//		fmt.Printf("uuid = %s, sha = %s\n", uuid, sha)
//		invokeArgs := make([][]byte, len(params))
//		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
//	}