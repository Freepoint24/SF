package main

import (
	"fmt"
)

const (
	operatorPlus           = "+"
	operatorMinus          = "-"
	operatorMultiplication = "*"
	operatorDivision       = "/"
)

func main() {
	var number1, number2, result float64
	var operator, operation string
	var err error

	fmt.Println("Введите первое число:")
	_, err = fmt.Scanln(&number1)
	if err != nil {
		fmt.Printf("ошибка при сканировании первого числа: %v", err)
	}

	fmt.Printf("Введите оператор(%s, %s, %s, %s):\n", operatorPlus, operatorMinus, operatorMultiplication, operatorDivision)
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Printf("ошибка при сканировании оператора: %v", err)
	}

	fmt.Println("Введите второе число:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Printf("ошибка при сканировании второго числа: %v", err)
	}

	if operator == operatorDivision && number2 == 0 {
		fmt.Println("ошибка: на ноль делить нельзя")
		return
	}

	switch operator {
	case operatorPlus:
		result = number1 + number2
		operation = "сложения"
	case operatorMinus:
		result = number1 - number2
		operation = "вычитания"
	case operatorMultiplication:
		result = number1 * number2
		operation = "умножения"
	case operatorDivision:
		result = number1 / number2
		operation = "деления"
	}

	fmt.Printf("Результат %s: %.2f", operation, result) //%.2f - ограничить количество знаков после запятой до двух (для чисел с плавающей точкой)
}
