package main

import (
	"app/electronic"
	"fmt"
)

func main() {
	f := func(b float64) float64 { return 0.0 }
	if f(5) > 0 {
		fmt.Println(f(5))
	}
	newIPhone := electronic.NewApplePhone("SE2")
	printCharacteristics(newIPhone)
}

func printCharacteristics(phone electronic.Phone) {

}
