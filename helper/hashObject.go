package helper

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashObject(data string) (string, error) {
	dataBytes := []byte(data)
	hash := sha1.New()
	hash.Write(dataBytes)
	oid := hex.EncodeToString(hash.Sum(nil))

	return oid, nil
}
