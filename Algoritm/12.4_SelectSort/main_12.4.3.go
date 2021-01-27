package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.4.3
//Реализуйте двунаправленную сортировку выбором.
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	//Реализуйте двунаправленную сортировку выбором.
	length := len(ar)
	for i := 0; i < (length - 1); i++ {
		for j := (length - 1); j > i; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			}
		}

	}
	//
	fmt.Println(ar)

}
