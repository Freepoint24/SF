package main

import "fmt"

func main()  {
	a := 3
	b := 2

	// Равно. Выведет false
	fmt.Println(a == b)
	// Не равно.  Выведет true
	fmt.Println(a != b)
	// Больше. Выведет true
	fmt.Println(a > b)
	// Больше или равно. Выведет true
	fmt.Println(a >=b)
	//Меньше. Выведет false
	fmt.Println(a < b)
	//Меньше ил равно. Выведет false
	fmt.Println(a <= b)
	
}