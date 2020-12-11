package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("anonimus")
	}()
	fmt.Println("string")
}
