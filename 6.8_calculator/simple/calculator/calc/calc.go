package calc

import (
	"fmt"
	"math"
)

type calculator struct {
	a string
	b string
	c int
}

//NewCalculator Creates a new instance of calculator
func NewCalculator() calculator {
	return calculator{}
}

func (c calculator) Calculate(num1, num2 float64, operator string) float64 {

	var result float64
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "%":
		result = mod(num1, num2)
	case "^":
		result = power(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		break
	}

	return result
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("Attempted vision by zero")
		return 0
	}
	return num1 / num2
}

func mod(num1, num2 float64) float64 {
	return math.Mod(num1, num2)
}

func power(num1, num2 float64) float64 {
	return math.Pow(num1, num2)
}

func (c *calculator) addition(num1, num2 float64) float64    { return num1 + num2 }
func (c *calculator) subtraction(num1, num2 float64) float64 { return num1 - num2 }
