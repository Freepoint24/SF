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
	printInterface(Comment{Text: "Comment massage"})
	printInterface(&Comment{Text: "comment pointer"})
	printInterface(&integ)
}

func printInterface(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("interface contains string value")
	case bool:
		fmt.Println("interface contains bool value")
	case Comment:
		comment := i.(Comment)
		fmt.Println(comment.Text)
	default:
		fmt.Println("unknown type")
	}
}
