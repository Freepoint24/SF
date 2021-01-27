package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.4.2
//Реализуйте сортировку выбором selectionSortByMax, работающую «справа налево»
//(поиск максимальных элементов и перемещение их в конец).
//https://play.golang.org/p/Hnhmz2nDhuE
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	selectionSortByMax(ar)

	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//Реализуйте сортировку выбором selectionSortByMax, работающую «справа налево»
//(поиск максимальных элементов и перемещение их в конец).
func selectionSortByMax(ar []int) {
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 0; i < len(ar); i++ {
		//minIndex := i
		maxIndex := i
		j := i + 1
		for j < len(ar) {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
			j = j + 1
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
		fmt.Printf("Sorting list:\t%v\n", ar) // визуализация поиска
	}

}
