package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Selection sort (Сортировка выбором)
//Задание 12.4.3
//Реализуйте двунаправленную сортировку выбором.
//https://play.golang.org/p/yQovKsYHLfh
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	selectionSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив

}

//Реализуйте двунаправленную сортировку выбором.
func selectionSort(ar []int) {
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 0; i < (len(ar) - 1); i++ {
		for j := (len(ar) - 1); j > i; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			}

		}
		fmt.Printf("Sorting list:\t%v\n", ar) // визуализация сортировки
	}
}

//варинат реализации
//https://play.golang.org/p/S8Yif4c-5A4
//	for i := 0; i < len(ar)-1; i++ {
//		min := i
//		for j := i + 1; j < len(ar); j++ {
//			if ar[min] > ar[j] {
//				min = j
//			}
//		}
//
//		if min != i {
//			swap(ar, i, min)
//		}
//
//	}
//}
//
//func swap(ar []int, i, j int) {
//	tmp := ar[i]
//	ar[i] = ar[j]
//	ar[j] = tmp
//	fmt.Printf("Sorting list:\t%v\n", ar) // визуализация сортировки
//}
