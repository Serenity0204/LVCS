package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

const GIT_DIR string = ".lvcs"

func Init() {
	os.Mkdir(GIT_DIR, 0755)
}

func hashObject(data string) (string, error) {
	dataBytes := []byte(data)
	hash := sha1.New()
	hash.Write(dataBytes)
	oid := hex.EncodeToString(hash.Sum(nil))

	return oid, nil
}

func main() {
	Init()
	// Example usage
	data := "Hello, World!"
	oid, err := hashObject(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Object ID:", oid)
	decodedString := string(data)
	fmt.Println("Decoded content:", decodedString)

	data = "Goodbye, World!"
	oid, err = hashObject(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Object ID:", oid)
	decodedString = string(data)
	fmt.Println("Decoded content:", decodedString)
}
