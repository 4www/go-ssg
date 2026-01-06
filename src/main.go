package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Building site:")
	file, err := os.ReadFile("./content/pages/index.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file", string(file))
}

