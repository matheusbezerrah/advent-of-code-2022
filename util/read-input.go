package util

import (
	"fmt"
	"os"
)

func ReadInput(path string) string {
	fmt.Println("Reading file...")
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}
