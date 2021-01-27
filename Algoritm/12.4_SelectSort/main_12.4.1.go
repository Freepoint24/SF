package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задание 12.4.1
//Реализуйте сортировку выбором, работающую «слева направо»
//(поиск минимальных элементов и перемещение их в начало).

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	selectionSort(ar)

	fmt.Println(ar)
}

// ваш код здесь
// Реализуйте сортировку выбором, работающую «слева направо»
//(поиск минимальных элементов и перемещение их в начало).

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
}
