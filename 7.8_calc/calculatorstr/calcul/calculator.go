package calc

import (
	"fmt"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

// блок вычислений
// приемник *calculator , название метода Calc , результат float64

func (a *calculator) Calc(output, number1, number2 float64, operator string) float64 {
	switch operator {
	case "*":
		output = number1 * number2
		fmt.Println("Результат умножения: output")
	case "+":
		output = number1 + number2
		fmt.Println("Результат сложения:")
	case "-":
		output = number1 - number2
		fmt.Println("Результат вычитания:")
	case "/":
		if number2 == 0 {
			fmt.Println("ошибка")
		} else {
			output = number1 / number2
			fmt.Println("Результат деления:")
		}
	default:
		fmt.Println("Ошибка")

	}
	return output

}
