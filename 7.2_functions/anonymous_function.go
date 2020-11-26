package main

import "fmt"

func main() {
	c := 4

	someFunction := func(x, y int) int { return x + y + c }

	result := someFunction(5, 6)

	fmt.Println(result)
}
