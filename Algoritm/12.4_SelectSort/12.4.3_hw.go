package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Selection sort (Сортировка выбором)
//Задание 12.4.3
//Реализуйте двунаправленную сортировку выбором.

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	//ar :=[]int{5, 3, 4, 1, 2}
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	selectionSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив

}

//Реализуйте двунаправленную сортировку выбором.
func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}
		if min != i {
			swap(ar, i, min)
		}
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
	fmt.Printf("Sorting list:\t%v\n", ar) // визуализация сортировки
}
