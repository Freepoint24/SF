package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2 float64
	var a, b, c float64

	// присвоить с 5 чтобы получить панику (D < 0)
	// присвоить а 4 чтобы получить равные корни (D = 0)
	a, b, c = 2, 4, 1

	discriminant := math.Pow(b, 2) - 4*a*c
	fmt.Printf("discriminant: %f\n", discriminant)

	if discriminant < 0 {
		panic("roots are complex values")
	}

	x1 = (-b + discriminant) / 2 * a
	x2 = (-b - discriminant) / 2 * a

	fmt.Printf("x1: %f, x2: %f", x1, x2)
}
