package main

import (
	"fmt"

	"github.com/Serenity0204/LVCS/helper"
)

func main() {
	helper.Init()
	// Example usage
	data := "Hello, World!"
	oid, err := helper.HashObject(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Object ID:", oid)
	decodedString := string(data)
	fmt.Println("Decoded content:", decodedString)

	data = "Goodbye, World!"
	oid, err = helper.HashObject(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Object ID:", oid)
	decodedString = string(data)
	fmt.Println("Decoded content:", decodedString)
}
