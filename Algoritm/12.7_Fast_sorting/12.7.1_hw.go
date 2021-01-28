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
	//ar := make([]int, 50)
	//for i := range ar {
	//	ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	ar := []int{9, 14, 5, -6, 0, 2, -8, 11}
	fmt.Printf("Unsorted list:\t%v\n", ar)
	quickSort(ar)
	fmt.Printf("Sorted list:\t%v\n", ar)

}

//quickSort(ar)
//fmt.Println(ar)
//}
func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	quickSort(ar[:split])
	quickSort(ar[split:])

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
