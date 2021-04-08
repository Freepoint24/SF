package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Insertion sort (Сортировка выбором)
//Задание 12.5.1
//Реализуйте сортировку вставками.
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

	fmt.Printf("Unsorted list:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	InsertionSort2(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//Insertion sort (Сортировка вставками).
func InsertionSort2(ar []int) {
	for i := 1; i < len(ar); i++ { //переменная i - количество проходов по слайсу
		x := ar[i]                         //переменная x - i-тый индекс массива
		j := i                             //переменная j присвоить значение i
		for ; j >= 1 && ar[j-1] > x; j-- { //если j больше или равно 1 и предыдущий элемент больше итого значения
			ar[j] = ar[j-1]
		}
		ar[j] = x
		fmt.Printf("Sorting ...:\t%v\n", ar) // визуализация
	}
}
