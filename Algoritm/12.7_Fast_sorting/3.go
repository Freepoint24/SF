package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Основная идея - берем опорную точку (pivot), проходим массив,
//чтобы элементы, которые меньше опорной точки оказались слева от нее, а которые больше - справа.
//Дальше берем часть массива до опорной точки и вторую часть после опортной точки, повторяем на них сортировку.
//Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.
//https://play.golang.org/p/cHcbiE0Ugry
func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar2 := make([]int, 50)
	for i := range ar2 {
		ar2[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	//fmt.Println(ar2)

	//func main() {
	//ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	//ar := ar2
	fmt.Printf("Unsorted list:\t%v\n", ar2)
	quicksort(ar2)
	//fmt.Printf("Sorted list:\t%v\n", ar2)

}

func quicksort(ar2 []int) {
	//fmt.Printf("Sorting ...:\t%v\n", ar)
	if len(ar2) <= 1 {
		return
	}

	split := partition(ar2)

	quicksort(ar2[:split])
	quicksort(ar2[split:])

}

func partition(ar2 []int) int {
	pivot := ar2[len(ar2)/2]

	left := 0
	right := len(ar2) - 1

	for {
		for ; ar2[left] < pivot; left++ {
		}

		for ; ar2[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar2, left, right)

	}
}
func swap(ar2 []int, i, j int) {
	tmp := ar2[i]
	ar2[i] = ar2[j]
	ar2[j] = tmp

}
