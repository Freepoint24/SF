package main

import "fmt"

type Comment struct {
	Text         string
	ChildComment *Comment
}

func main() {
	com1 := &Comment{Text: "com1"}

	com1Sub := &Comment{Text: "cjm1Sub"}
	com1.ChildComment = com1Sub

	com1SubSub := &Comment{Text: "com1SubSub"}
	com1Sub.ChildComment = com1SubSub

	printRecursive(com1)
}

/*com1
com1Sub
	com1SubSub
*/

func printRecursive(c *Comment) {
	fmt.Println(c.Text)

	if c.ChildComment != nil {
		printRecursive(c.ChildComment)
	}
}
