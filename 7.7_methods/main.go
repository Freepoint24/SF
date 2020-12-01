package main

import "fmt"

type A struct {
	i int
}

func (a A) square() int {
	return a.i * a.i
}

func main() {
	a := A{i: 10}
	a1 := A{i: 5}

	fmt.Println(a.square())
	fmt.Println(a1.square())
}
