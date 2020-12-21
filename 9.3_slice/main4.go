package main

import "fmt"

func main()  {
	s := make([]int, 3, 4)

	fmt.Println(s)

	example(s)

	fmt.Println(s)
	}

func example(s []int)  {
	s[0] = 100
}
