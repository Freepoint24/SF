package main

import "fmt"

func main() {
	someVariable := "some text"

	ourCallback := func() {
		someVariable = "other string"

		fmt.Println(someVariable)
	}
	runCallback(ourCallback)

	// выведет "other string"
	fmt.Println(someVariable)
}

func runCallback(callback func()) {
	callback()

}
