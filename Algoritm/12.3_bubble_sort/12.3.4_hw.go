package main

import (
	"fmt"
	"math/rand"
	"time"
)

//алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по убыванию.
//https://play.golang.org/p/12v0Ozq-Qla
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	bubbleSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//ваш код здесь
func bubbleSort(ar []int) {
	// sort on the left - по убыванию
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 0; i < (len(ar) - 1); i++ {
		for j := 0; j < ((len(ar) - 1) - i); j++ {
			if ar[j] < ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
		//fmt.Printf("Sorting ...:\t%v\n", ar) //визуальное представление сортировки
	}
}

//вариант реализации
//https://play.golang.org/p/tPE0T6_DXka
//for i := 0; i < len(ar); i++ {
//for j := len(ar) - 1; j > i; j-- {
//if ar[j-1] < ar[j] {
//swap(ar, j-1, j)
//			}
//		}
//		}
//}
//func swap(ar []int, i, j int) {
//	tmp := ar[i]
//	ar[i] = ar[j]
//	ar[j] = tmp
//fmt.Printf("Sorting ...:\t%v\n", ar) //визуальное представление сортировки
//}
