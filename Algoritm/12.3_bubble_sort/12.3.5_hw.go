package main

import (
	"fmt"
	"math/rand"
	"time"
)

//спорное решение
//алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию,
//используя рекурсивные функции.
//https://play.golang.org/p/dCGscEpjv42
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
	// sort on the left - по возрастанию
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				swap(ar, j-1, j)
			}
		}
	}
}
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
	fmt.Printf("Sorting ...:\t%v\n", ar) //визуальное представление сортировки

}
