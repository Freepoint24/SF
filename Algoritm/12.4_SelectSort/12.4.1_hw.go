package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.4.1
//Реализуйте сортировку выбором Selection sort, работающую «слева направо»
//(поиск минимальных элементов и перемещение их в начало).

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

// ваш код здесь
// Реализуйте сортировку выбором, работающую «слева направо»
//(поиск минимальных элементов и перемещение их в начало).

func selectionSort(ar []int) {
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 0; i < len(ar); i++ {
		minIndex := i
		j := i + 1
		for j < len(ar) {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			j = j + 1
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
		fmt.Printf("Sorting list:\t%v\n", ar) // визуализация сортировки
	}

}

////вариант реализации
////https://play.golang.org/p/mxIXCXf7pUc
//	for i := 0; i < len(ar)-1; i++ {
//		min := i
//		for j := i + 1; j < len(ar); j++ {
//			if ar[min] > ar[j] {
//				min = j
//			}
//		}
//		if min != i {
//			swap(ar, i, min)
//		}
//	}
//}
//
//func swap(ar []int, i, j int) {
//	tmp := ar[i]
//	ar[i] = ar[j]
//	ar[j] = tmp
//	fmt.Printf("Sorting ...:\t%v\n", ar) //визуальное представление сортировки
//}
