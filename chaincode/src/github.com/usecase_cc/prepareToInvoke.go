package main
	
import (
	"crypto/sha256"
	"encoding/base64"
)

func prepareToInvoke(uuid string, logString string) [][]byte{
	h := sha256.New()
	h.Write([]byte(logString))
	b := h.Sum(nil)
	base64Enc := base64.StdEncoding.EncodeToString(b)
	params := []string{"set", uuid, base64Enc}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {invokeArgs[i] = []byte(arg)}
	
	return invokeArgs 
}
