package main

import "fmt"

//алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
//учебный пример play.golang.org/p/jwIZN_WL8sg
func bubbleSort(ar []int) {
	for i := 0; i < (len(ar) - 1); i++ {
		for j := 0; j < ((len(ar) - 1) - i); j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}

//
//Задание 12.4.1
//Реализуйте сортировку выбором Selection sort, работающую «слева направо»
//(поиск минимальных элементов и перемещение их в начало).
func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		minIndex := i
		j := i + 1
		for j < len(ar) {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			j = j + 1
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

//
//Selection sort (Сортировка выбором)
//Задание 12.4.3
//Реализуйте двунаправленную сортировку выбором.
func selectionSort(ar []int) {
	for i := 0; i < (len(ar) - 1); i++ {
		for j := (len(ar) - 1); j > i; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			}

		}
	}
}

//
//Insertion sort (Сортировка выбором)
//Задание 12.5.1
//Реализуйте сортировку вставками.
//Реализуйте двунаправленную сортировку выбором.
func insertionSort(ar []int) {

	for i := 1; i < len(ar); i++ {
		for j := i; j != 0 && ar[j] < ar[j-1]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
		fmt.Printf("Sorting ...:\t%v\n", ar) // визуализация
	}
}
