//Сложность bubbleSort
//В лучшем случае: O(n).
//В среднем случае: O(n2).
//В худшем случае: O(n2) Е.
//Ёмкостная, в худшем: O(1).

//12.3.2 алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
//учебный пример
package main

import (
	"fmt"
	//"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := []int{5, 3, 4, 1, 2} // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	bubbleSort(s1)

}

func bubbleSort(s1 []int) {
	length := len(s1)                   //переменная length - длина слайса s1
	for i := 0; i < (length - 1); i++ { // переменная i - количество проходов по слайсу
		for j := 0; j < ((length - 1) - i); j++ { //переменая j - индекс элемента слайса
			if s1[j] > s1[j+1] { //если предыдущее значение меньше следующего, то
				s1[j], s1[j+1] = s1[j+1], s1[j] //поменять их местами
			}
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
