package main
	
import (
	"strings"
	"crypto/sha256"
	"encoding/base64"
)

func prepareToInvoke(uuid string, logString string) [][]byte{
	logStringWoutN := strings.TrimSuffix(logString, "\n")
	h := sha256.New()
	h.Write([]byte(logStringWoutN))
	b := h.Sum(nil)
	base64Enc := base64.StdEncoding.EncodeToString(b)
	params := []string{"set", uuid, base64Enc}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {invokeArgs[i] = []byte(arg)}

	return invokeArgs 
}
