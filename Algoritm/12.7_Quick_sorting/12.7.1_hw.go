package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.7.1
//Реализуйте быструю сортировку.
//https://play.golang.org/p/ltbNwisNTFg

var index, start, end int

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar :=[]int{6,5, 3, 1, 8, 7, 2, 4} // пример из задания
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	//ar := make([]int, 50)
	//for i := range ar {
	//	ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	//}
	fmt.Printf("Unsorted list:\t%v\n", ar)
	//fmt.Println("")
	quickSort(ar, 0, len(ar)-1)
	//fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar)
}

// ваш код здесь
func quickSort(ar []int, start, end int) {
	if start >= end {
		return
	}

	pivot := ar[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > ar[j] {
			ar[i], ar[j] = ar[j], ar[i]
			i++
		}
		fmt.Printf("Sorting ...:\t%v\n", ar) // визуализация
	}

	ar[start], ar[i-1] = ar[i-1], ar[start]

	quickSort(ar, start, i-2)
	quickSort(ar, i, end)
}
