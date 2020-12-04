package calc

import (
	"fmt"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

// блок вычислений
func (a *calculator) Calc(output, number1, number2 float64, operator string) float64 {

	switch operator {
	case "*":
		output = number1 * number2
		fmt.Println("Результат умножения:", NewCalculator())
	case "+":
		output = number1 + number2
		fmt.Println("Результат сложения:", output)
	case "-":
		output = number1 - number2
		fmt.Println("Результат вычитания:", output)
	case "/":
		if number2 == 0 {
			fmt.Println("ошибка")
		} else {
			output = number1 / number2
			fmt.Println("Результат деления:", output)
		}
	default:
		fmt.Println("Ошибка")
	}
}
