//package main

import (
	"fmt"
)

//Основная идея - берем опорную точку (pivot), проходим массив,
//чтобы элементы, которые меньше опорной точки оказались слева от нее, а которые больше - справа.
//Дальше берем часть массива до опорной точки и вторую часть после опортной точки, повторяем на них сортировку.
//Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.
//https://play.golang.org/p/cHcbiE0Ugry
func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Printf("Unsorted list:\t%v\n", ar)
	Quicksort(ar)
	fmt.Printf("Sorted list:\t%v\n", ar)

}

func Quicksort(ar []int) {
	fmt.Printf("Sorting ...:\t%v\n", ar)
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])

}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

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

}
