package calc

import (
	"fmt"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (a *calculator) Сalc(number1, number2, output float64, operator string) float64 {

	// блок вычислений
	switch operator {
	case "*":
		output = number1 * number2
		fmt.Println("Результат умножения:", output)
	case "/":
		if number2 == 0 {
			fmt.Println("ошибка")
		} else {
			output = number1 / number2
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
