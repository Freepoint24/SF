package main

import "fmt"

func main() {
	var val interface{} = 10

	number := val.(int)

	fmt.Printf("number: %d", number)
}
