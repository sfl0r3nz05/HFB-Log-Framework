package main
	
import (
	"crypto/sha256"
	"encoding/base64"
)

func prepareToInvoke(uuid string, logString string, CHANNEL_ENV string) [][]byte{
	hasher := sha256.New()
	hasher.Write([]byte(logString))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	params := []string{"set", uuid, sha}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {invokeArgs[i] = []byte(arg)}
	
	return invokeArgs 
}
