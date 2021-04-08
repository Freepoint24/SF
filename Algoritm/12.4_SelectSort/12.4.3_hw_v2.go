package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Selection sort (Сортировка выбором)
//Задание 12.4.3
//Реализуйте двунаправленную сортировку выбором.

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	//ar :=[]int{5, 3, 4, 1, 2}
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	ar := make([]int, 30)
	for i := range ar {
		ar[i] = rand.Intn(50) - 25 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	selectionSort2(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив

}

//Реализуйте двунаправленную сортировку выбором.
func selectionSort2(ar []int) {
	length := len(ar)                   //переменная lenght - длина слайса
	for i := 0; i < (length - 1); i++ { // переменная i - количество проходов слайсов
		for j := (length - 1); j > i; j-- { // переменная j - индекс слайса
			if ar[j] < ar[j-1] { //есди значение текущее меншье предыдущего, то
				ar[j], ar[j-1] = ar[j-1], ar[j] //поменять местами
			}
		}
		fmt.Printf("Sorting ...:\t%v\n", ar) // визуализация
	}
}
