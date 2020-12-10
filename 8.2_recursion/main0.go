package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}
