package calc

import (
	"log"
)

type calculator struct {}

func NewCalculator() calculator  {
	return calculator{}
}

func (c *calculator) Calculate(number1, number2 float64, operator string) float64 {
	switch operator {
	case "+":
			return c.add(number1, number2)
	default:
		log.Printf("ошибка")
		return 0
	}

//	return 0
}

func (c *calculator) add(number1, number2 float64) float64 {
	return number1 + number2
}