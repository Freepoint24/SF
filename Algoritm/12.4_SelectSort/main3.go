package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	//selectionSort(ar)
	fmt.Println(ar)

	//s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", ar)
	fmt.Println("")
	length := len(ar)
	for i := 0; i < (length - 1); i++ {
		for j := (length - 1); j > i; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			}
		}
		fmt.Printf("Sorting ...:\t%v\n", ar)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar)

}
