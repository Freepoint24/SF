package main

import "fmt"

func main() {
	i := 2

	if i%2 == 0 {
		fmt.Println("Делится на 2 без остаттка")
	} else if i%3 == 0 {
		fmt.Println("Делится на 3 без остатка")
	} else if i%5 == 0 {
		fmt.Println("делится на 5 без остатка")
	} else {
		fmt.Println("делится на что-то другое")
	}
}
