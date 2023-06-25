package main

import (
	"fmt"
	"os"
)

const GIT_DIR string = ".lvcs"

func Init() {
	os.Mkdir(GIT_DIR, 0755)
}

func main() {
	Init()
	fmt.Println("Hello")
}
