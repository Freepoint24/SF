package main

import (
	calc "calculator-b/calcul"
	"fmt"
)

// Блок ввода данных
func main() {
	var number1, number2 float64
	var operator string
	var err error

	//ввод числа 1
	fmt.Print("Введите число 1:")
	_, err = fmt.Scanln(&number1)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}

	//ввод оператора
	fmt.Print("Введите оператор:")
	_, _ = fmt.Scanln(&operator)
	switch operator {
	case "+":
		operator = "+"
	case "-":
		operator = "-"
	case "*":
		operator = "*"
	case "/":
		operator = "/"
	default:
		fmt.Printf("Ошибка ввода оператора")
		return
	}

	//ввод числа 2
	fmt.Print("Введите число 2:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Printf("ошибка при вводе второго числа")
		return
	}

	fmt.Println(calc.NewCalculator())
}
