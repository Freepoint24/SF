package main

import (
	"fmt"
)

var num1, num2, result float32
var literal string

func main() {
	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println(fmt.Sprintf("Вводите только числа! %s", err))
		return
	}
	{
		fmt.Println("Введите литерал: ")
		_, _ = fmt.Scanln(&literal)
		switch literal {
		case "+":
			literal = "+"
		case "-":
			literal = "-"
		case "*":
			literal = "*"
		case "/":
			literal = "/"
		default:
			fmt.Println("Вводите только '+', '-', '*', '/'")
			return
		}
	}
	{
		fmt.Println("Введите второе число: ")
		_, err := fmt.Scanln(&num2)
		if err != nil {
			fmt.Println(fmt.Sprintf("Вводите только числа! %s", err))
			return
		}
	}
	if literal == "/" && num2 == 0 {
		fmt.Println("Делить на 0 нельзя!")
		return
	}
	{
		switch literal {
		case "+":
			result = num1 + num2
			println("Результат сложения: ", result)
		case "-":
			result = num1 - num2
			println("Результат вычитания: ", result)
		case "*":
			result = num1 * num2
			println("Результат умножения: ", result)
		case "/":
			result = num1 / num2
			println("Результат деления: ", result)
		}
	}
}
