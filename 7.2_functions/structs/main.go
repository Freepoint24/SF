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
	b := B{a: a}
	c := C{b: b}

	fmt.Println(a.i)
	fmt.Println(b.a.i)
	fmt.Println(c.b.a.i)
}
