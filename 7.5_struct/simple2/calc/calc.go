package calc

import "fmt"

type calculator struct {
	x float64
	y float64
	r float64
}

func NewCalc() {
	c := calculator{5, 7, 8}

	fmt.Println(c)
}
