package models

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
