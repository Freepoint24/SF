package main2

import "fmt"

func main()  {
	s := make([]int, 1, 2)

	fmt.Println(s, len(s), cap(s), fmt.Sprintf("%p",&s))

	s = append(s, 10, 20)


	fmt.Println(s, len(s), cap(s), fmt.Sprintf("%p", &s))
}
