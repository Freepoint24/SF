package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	x1, x2, err := squareRoot(2, 4, 1)

	if err != nil {
		panic(err)
	}

	fmt.Printf("x1: %f, x2: %f", x1, x2)
}

func squareRoot(a, b, c float64) (float64, float64, error) {
	// Функция для вычисления корней квадратного уравнения
	var x1, x2 float64

	discrimint := math.Pow(b, 2) - 4*a*c
	fmt.Sprintf("discriminant: %f\n", discrimint)

	if discrimint < 0 {
		return 0, 0, errors.New("roots are complex values")
	}

	x1 = (-b + discrimint) / 2 * a
	x2 = (-b - discrimint) / 2 * a

	return x1, x2, nil
	// Функция для вычисления корней квадратного уравнения
}
