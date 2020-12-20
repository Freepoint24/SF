package main2

import "fmt"

func main()  {
	x := [2][2]int{
	[2]int{1,2},
	[2]int{3,4},
	}

	for i, arr := range x {
		fmt.Println(i, "=>", arr)
	}
}
