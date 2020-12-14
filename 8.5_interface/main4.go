package main

import "fmt"

type Comment struct {
	Text string
}

func main() {
	integ := 7

	printInterface(10)
	printInterface("str")
	printInterface(true)
	printInterface(Comment{Text: "Comment"})
	printInterface(&Comment{Text: "comment pointer"})
	printInterface(&integ)
}

func printInterface(i interface{}) {
	fmt.Println(i)
}
