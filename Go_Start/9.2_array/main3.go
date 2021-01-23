package main2

import "fmt"

func main() {
	var x = [...]int{10, 11, 12, 13, 14, 15}


	for i, v := range x {
		fmt.Println(i, "=>", v)

	}

}

