package main

import "fmt"

func main()  {

	 var c int32

	fmt.Println("Введите число")
	fmt.Scan(&c)

	switch {

	case c == 0:
		fmt.Println("Ноль")
	case c > 0:
		fmt.Println("Число положительное")
	case c < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("ошибка")
	}

}