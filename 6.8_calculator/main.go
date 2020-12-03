package main

import (
	"fmt"
)
func main() {
	var number1, number2, output float32
	var operator string
	var err error

	//ввод числа 1
	fmt.Print("Введите число 1:")
	//fmt.Scan(&number1)
	_, err = fmt.Scanln(&number1)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}
	{
	//ввод оператора
		fmt.Print("Введите оператор:")
		_, _ = fmt.Scanln(&operator)
		switch operator {
		case "+": operator = "+"
		case "-": operator = "-"
		case "*": operator = "*"
		case "/": operator = "/"
		default:
			fmt.Printf("Ошибка ввода оператора")
			return
		}
	}
	{
	//ввод числа 2
		fmt.Print("Введите число 2:")
		_, err = fmt.Scanln(&number2)
		if err != nil {
			fmt.Printf("ошибка при вводе второго числа")
			return
		}

	// блок вычислений
		switch operator {
		case "*":
			output = number1 * number2
			fmt.Println("Результат умножения:", output)
		case "/":
			if number2 == 0 {
				fmt.Println("ошибка")
			} else {
				output = number2 / number1
				fmt.Println("Результат деления:", output)
			}
		case "+":
			output = number1 + number2
			fmt.Println("Результат сложения:", output)
		case "-":
			output = number1 - number2
			fmt.Println("Результат вычитания:", output)
		default:
			fmt.Println("Ошибка")
		}

	}
}