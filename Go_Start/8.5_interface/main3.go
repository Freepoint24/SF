package main

import "fmt"

type Book struct {
	Author string
	Title  string
}

func (b *Book) String() string {
	return fmt.Sprintf("Book '%s' written %s", b.Title, b.Author)
}

type Magazine struct {
	Title  string
	Number int
}

func (m *Magazine) String() string {
	return fmt.Sprintf("Magazin '%s, number: %d", m.Title, m.Number)
}

func main() {
	b := Book{Title: "Harry Potter", Author: "Joan Rouling"}
	m := Magazine{Title: "Top gear", Number: 32}

	printStringerInterface(&b)
	printStringerInterface(&m)
}

func printStringerInterface(i fmt.Stringer) {
	switch converted := i.(type) {
	case *Book:
		fmt.Println(converted.Author)
	case *Magazine:
		fmt.Println(converted.Number)
	}
}
