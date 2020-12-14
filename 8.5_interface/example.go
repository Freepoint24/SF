package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

func main() {
	a := Article{
		Title:  "Основы программирования на GO",
		Author: "James Bond",
	}
	b := Book{
		Title:  "Harry Potter",
		Author: "Joanne Rowling",
	}

	// Print(a)
	typeCast(a)
	typeCast(b)
}

func typeCast(s Stringer) {
	switch s.(type) {
	case Article:
		a := s.(Article)
		fmt.Println(a)
	case Book:
		b := s.(Book)
		fmt.Println(b)
	}
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

type Stringer interface {
	String() string
}
