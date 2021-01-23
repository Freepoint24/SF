package main

import "fmt"

func main()  {
	s := make([]int, 3, 4)

	fmt.Println(s)

	example(s)

	fmt.Println(s)
	}

func example(s []int)  {
	sCopy := make([]int, len(s))

	copy(sCopy, s)

	sCopy[0] = 100
	fmt.Println(sCopy)

}
