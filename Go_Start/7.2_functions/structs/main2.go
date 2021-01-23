package main

import "fmt"

type A struct {
	i int
}
type B struct {
	a A
}
type C struct {
	b B
}

func main() {
	a := A{i: 10}

	exampleStruct(&a)
	// &
	fmt.Println(a.i)
}

func exampleStruct(exanpleStruct *A) {
	exanpleStruct.i = exanpleStruct.i + 100
}
