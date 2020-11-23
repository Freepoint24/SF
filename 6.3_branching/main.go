package main

import "fmt"

func main() {
	i := 5

	// Выполнится только первое истинное условие
	if i > 0 {
		fmt.Println(0)
	} else if i > 1 {
		fmt.Println(1)
	} else if i > 2 {
		fmt.Println(2)
	}
	// Выполнстя все условия соответствующие истинности
	if i > 0 {
		fmt.Println(0)
	}
	if i > 1 {
		fmt.Println(1)
	}
	if i > 2 {
		fmt.Println(2)
	}
}