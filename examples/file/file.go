package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("file.json")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Content: \n%s\n", string(content))
	}

	content, err = os.ReadFile("unknown.json")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Content: \n%s\n", string(content))
	}
}
