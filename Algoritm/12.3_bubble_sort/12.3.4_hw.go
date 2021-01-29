package main

import (
	"fmt"
	"math/rand"
	"time"
)

//12.3.4 алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по убыванию.
//https://play.golang.org/p/ICaHAehXedA
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	//ar :=[]int{6, 5, 3, 1, 8, 7, 2, 4} //пример из учебной главы
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	bubbleSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//ваш код здесь sort on the right - по убыванию
func bubbleSort(ar []int) {
	for i := 0; i < (len(ar) - 1); i++ {
		for j := 0; j < ((len(ar) - 1) - i); j++ {
			if ar[j] < ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
		fmt.Printf("Sorting ...:\t%v\n", ar) //визуальное представление сортировки
	}
}
