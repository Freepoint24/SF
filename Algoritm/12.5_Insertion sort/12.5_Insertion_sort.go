package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Insertion sort (Сортировка выбором)
//Задание 12.5.1
//Реализуйте сортировку вставками.
//https://play.golang.org/p/OHWF2q482al
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	insertionSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//Реализуйте двунаправленную сортировку выбором.
func insertionSort(ar []int) {
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 1; i < len(ar); i++ {
		for j := i; j != 0 && ar[j] < ar[j-1]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
		//fmt.Printf("Sorting ...:\t%v\n", ar)// визуализация
	}
}

//вариант реализации
//	for i := 1; i < len(ar); i++ {
//		x := ar[i]
//		j := i
//		for ; j >= 1 && ar[j-1] > x; j-- {
//		ar[j] = ar[j-1]
//
//		}
//	ar[j] = x
//		fmt.Printf("Sorting ...:\t%v\n", ar)// визуализация
//	}
//}
