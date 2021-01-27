package main

import (
	"fmt"
	"math/rand"
	"time"
)

//алгоритм сортировки пузырьком, который отсортирует массив int по возрастанию.
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	bubbleSort(ar)

	fmt.Println(ar)
}

//ваш код здесь
func bubbleSort(ar []int) {
	//		if ar[i] > ar[j] {
	//			swap(ar, i, j)
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

}
