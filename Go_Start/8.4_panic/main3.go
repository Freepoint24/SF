package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(deferExample())
}

func deferExample() error {
	defer func() {
		fmt.Println("working")
	}()
	return errors.New("error msg")
}
