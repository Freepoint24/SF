package main

import "fmt"

func main() {
	i := 5
	for i <= 10 {
		fmt.Println(i)
		//i = i + 1
		i++
	}
	fmt.Println("Program finished")
}
