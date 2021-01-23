package main

import "fmt"

func main() {
	someVariable := "some text"

	runCallback(func() {
		someVariable = "other string"
		fmt.Println(someVariable)
	})
	// выведет "other string"
	fmt.Println(someVariable)
}

func runCallback(callback func()) {
	callback()
}
