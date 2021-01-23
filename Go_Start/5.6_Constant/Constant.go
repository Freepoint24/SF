package main

import "fmt"

const pi = 3.14

const (
	One = iota
	Two
	Three
)

func main() {
	const x string = "Hello World"
	println(x)

	fmt.Println("Математическая константа PI =", pi)

	// pi = 30

	// Выведет 0 1 2, потмоу что iota начинается с 0
	fmt.Println(One, Two, Three)

}
