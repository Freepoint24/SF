package main

import "fmt"

func main()  {
	arr := [3]int{10, 20, 30}
	s := arr[0:1]

	fmt.Println(arr)
	fmt.Println(s)

	arr[0] = 100

	fmt.Println(arr[0])
	fmt.Println(s[0])
}main2.go
