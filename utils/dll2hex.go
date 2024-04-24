package utils

import (
	"encoding/hex"
	"io/ioutil"
)

func BinToHex(DllPath string) string {
	bytes, err := ioutil.ReadFile(DllPath)
	if err != nil {
		return ""
	}
	DLLHex := hex.EncodeToString(bytes)
	return DLLHex
}
