// Почему именно существует параметр CAPACITY при создании среза в Golang

package main

import "fmt"

func main() {

	fmt.Println("With 0 capacity")
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(i, &s[0])
	}

	fmt.Println("With 10 capacity")
	s = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(i, &s[0])
	}

}
