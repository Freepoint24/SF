package main

import "fmt"

func main()  {
	// Присвоение
	a := 5

	// a = a + 1 Присвоение со сложением
	a += 1
	fmt.Println(a)

	// a = a - 1 Присвоение с вычитанием
	a -= 1
	fmt.Println(a)

	// a = a * 2 Присвоение с умноженим
	a *= 2
	fmt.Println(a)

	// a = a / 2 Присвоение с деление
	a /= 2
	fmt.Println(a)

	// a = a % 2 Присвоение с получением остатка от деления
	a %= 2
	fmt.Println(a)

}
