package main

import "fmt"

func main() {
	defer one()
	defer two()
	defer three()
}

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}
