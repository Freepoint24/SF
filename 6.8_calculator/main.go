package main

import "fmt"

func main() {
	var number1, number2, output float32
	var operator string

	//ввод числа 1
	fmt.Print("Введите число 1:")
	fmt.Scan(&number1)

	//ввод оператора
	fmt.Print("Введите оператор:")
	fmt.Scan(&operator)

	//ввод числа 1
	fmt.Print("Введите число 2:")
	fmt.Scan(&number2)

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
