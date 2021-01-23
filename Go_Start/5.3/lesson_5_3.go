package main

import "fmt"

// Объявление переменной в области видимости всего пакета Lesson_5_3
// Любая функция внутри данного пакета может прочитать или зменить значение данной переменной
//Переменная с видна всему пакету
var c string = "Hello"

//Объявление нескольких переменных в одном блоке. Удобно для группировки близких по смыслу переменных
var (
	a = 5
	b = 10
	r = 15
)

// Функция main - точка входа в прогрмму. При запуске программы, выполнение начинается с нее
func main() {
	// Обяъвлене локальной перемнной внутри функции
	var x string = "Hello World"
	// Переменные видны в области видиомости функции main
	// Присвоение значения переменной
	s := "First"
	fmt.Println(s)
	s = "second"
	fmt.Println(s)

	//Краткое объявление переменной
	name := 10
	Y := "Max"

	fmt.Println("4", name, Y)
	fmt.Println(x)
	fmt.Println(Y, b)
}

// Функция f, которая так же имеет доступ к переменной 'c'
func g() {
	fmt.Println(c)
}
