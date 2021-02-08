package main

import (
	"fmt"
)

//1. строку преобразовать в массив
//2. массив в число
//3. найти остаток от деленя на 1000

func Pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	x := "abcdERAA" // строка для преобразования
	y := []byte(x)  //преобразовать строку в массив байт
	fmt.Println(x)
	fmt.Println(y)

	//2. массив в число
	nums := []int{1, 2, 3, 7, 8}
	num := 0
	length := len(nums)
	for i, d := range nums {
		num += d * Pow(10, length-i-1)
	}
	fmt.Println(num)
	//3. найти остаток от деленя на 1000

	z := num % 1000

	fmt.Println(z)

}
