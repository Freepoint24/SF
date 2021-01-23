package main

import "fmt"

func main() {
	c := 4

	sum := func() { c = c + 10 }

	example(sum)

	fmt.Println(c)
}

func example(callback func()) {
	callback()
}

// функция с переменным числом аргументов
func variadic(numbers ...int) {

}
