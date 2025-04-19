package httpx

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateETag(data []byte) string {
	hash := sha1.Sum(data)
	return `"` + hex.EncodeToString(hash[:]) + `"`
}
