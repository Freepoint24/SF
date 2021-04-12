package main

import "fmt"

//12.3.2
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
//Задание 12.4.1func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
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
//Задание 12.4.3 -
//Selection sort (Сортировка выбором)
//Реализуйте двунаправленную сортировку выбором.
//func selectionSort(ar []int) {
//	for i := 0; i < (len(ar) - 1); i++ {
//		for j := (len(ar) - 1); j > i; j-- {
//			if ar[j] < ar[j-1] {
//				ar[j], ar[j-1] = ar[j-1], ar[j]
//			}
//
//		}
//	}
//}

//Задание 12.5.1
//Insertion sort (Сортировка выбором)
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

//ЗАДАНИЕ 12.6.1
//Реализуйте сортировку слиянием, Merge sort
//https://play.golang.org/p/UiNZ2qSRdhd
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
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

//
//Задание 12.7.1
//Реализуйте быструю сортировку.
//https://play.golang.org/p/ltbNwisNTFg
// ваш код здесь
func quickSort(ar []int, start, end int) {
	if start >= end {
		return
	}

	pivot := ar[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > ar[j] {
			ar[i], ar[j] = ar[j], ar[i]
			i++
		}
	}

	ar[start], ar[i-1] = ar[i-1], ar[start]

	quickSort(ar, start, i-2)
	quickSort(ar, i, end)
}

//
