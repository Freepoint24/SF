/*Сложность
В лучшем случае: O(n log(n)).
В среднем случае: O(n log(n)).
В худшем случае: O(n log(n)).
Емкостная, в худшем: O(n).
*/
package main

import (
	"fmt"
	"math/rand"
	//"testing"
	"time"
)

//ЗАДАНИЕ 12.6.1
//Реализуйте сортировку слиянием, Merge sort
//https://play.golang.org/p/UiNZ2qSRdhd

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := []int{6, 5, 3, 1, 8, 7, 2, 4} // пример из задания
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	//ar := make([]int, 25)
	//for i := range ar {
	//	ar[i] = rand.Intn(100) - 50 // ограничиваем случайно значение от [-100;100]
	//}

	fmt.Printf("Unsorted list:\t%v\n", ar) // не сортированный массив
	sorted := mergeSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", sorted) // сортированный массив
}

// ваш код здесь
func mergeSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				fmt.Printf("Sorting:\t%v\n", result) //визауализация 1 из 4
				left = left[1:]
			} else {
				result = append(result, right[0])
				fmt.Printf("Sorting:\t%v\n", result) // визаулизация 2 из 4
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			fmt.Printf("Sorting:\t%v\n", result) // визуализация 3 из 4
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			fmt.Printf("Sorting:\t%v\n", result) // визуализация 4 из 4
			right = right[1:]
		}
	}

	return result
}
