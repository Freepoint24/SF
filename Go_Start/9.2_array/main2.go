package main2

import "fmt"

func main() {
	var x [4]int

	x[0] = 10
	x[1] = 11
	x[2] = 12
	x[3] = 13

	for i, v := range x {
		fmt.Println(i, "=>", v)

	}

}