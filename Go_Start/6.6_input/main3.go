package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите число:")

	fmt.Scanln(&number)

	fmt.Println(number)

}
