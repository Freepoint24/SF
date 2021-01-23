package main

import "fmt"

func main() {
	fmt.Println("before panic")

	panic("panic massage")

	fmt.Println("after panic")

}
