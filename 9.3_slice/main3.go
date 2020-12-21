package main2

import "fmt"

func main()  {
	s := make([]int, 3, 4)

	subS := s[:]

	fmt.Println(&s[0])
	fmt.Println(&subS[0])
}
