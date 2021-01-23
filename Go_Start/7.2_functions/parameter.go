package main

import "fmt"

func main() {
	someVariable := "some text"

	ourCallback := func() { fmt.Println(someVariable) }

	runCallback(ourCallback)

}

func runCallback(callback func()) {
	callback()
}
