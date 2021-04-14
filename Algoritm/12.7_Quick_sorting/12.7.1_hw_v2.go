package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.7.1
//Реализуйте быструю сортировку.
//https://play.golang.org/p/ltbNwisNTFg

//var index, start, end int

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := []int{6, 5, 3, 1, 8, 7, 2, 4} // пример из задания
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	//ar := make([]int, 50)
	//for i := range ar {
	//	ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	//}
	fmt.Printf("Unsorted list:\t%v\n", ar)
	//fmt.Println("")
	quicksort(ar)
	//fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar)
}

// ваш код здесь

func quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)
	fmt.Println("Опорный ", ar)
	quicksort(ar[:split])
	quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]
	fmt.Println("Опорный элемент ", pivot)
	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
	fmt.Println("Сортировка....:", ar)
}
