package main

import (
	"math/rand"
	"testing"
)

//запускать командой go test -bench=. ./...

//12.3.2
//алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
func Benchmark10BubbleSort(b *testing.B) {
	b.SetBytes(10)
}
func Benchmark1000BubbleSort(b *testing.B) {
	b.SetBytes(1000)
}
func Benchmark1000000BubbleSort(b *testing.B) {
	b.SetBytes(100000)
}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

//Задание 12.4.1
//Реализуйте сортировку выбором Selection sort, работающую «слева направо»
func Benchmark10SelectionSort(b *testing.B) {
	b.SetBytes(10)
}
func Benchmark1000SelectionSort(b *testing.B) {
	b.SetBytes(1000)
}
func Benchmark100000SelectionSort(b *testing.B) {
	b.SetBytes(100000)
}

func BenchmarkSelectionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

//Задание 12.4.3
//Selection sort (Сортировка выбором)
//func BenchmarkSelectionSort(b *testing.B) {
//	b.SetBytes(10)
//}

//Задание 12.5.1
//Insertion sort (Сортировка выбором)
func BenchmarkInsertionSort(b *testing.B) {
	b.SetBytes(10)
}
func Benchmark1000InsertionSort(b *testing.B) {
	b.SetBytes(1000)
}
func Benchmark100000InsertionSort(b *testing.B) {
	b.SetBytes(100000)
}

//ЗАДАНИЕ 12.6.1
//Реализуйте сортировку слиянием, Merge sort
func BenchmarkMergeSort(b *testing.B) {
	b.SetBytes(10)
}
func Benchmark1000MergeSort(b *testing.B) {
	b.SetBytes(1000)
}
func Benchmark100000MergeSort(b *testing.B) {
	b.SetBytes(100000)
}

//Задание 12.7.1
//Реализуйте быструю сортировку.
func BenchmarkQuickSort(b *testing.B) {
	b.SetBytes(10)
}
func Benchmark1000QuickSort(b *testing.B) {
	b.SetBytes(1000)
}
func Benchmark100000QuickSort(b *testing.B) {
	b.SetBytes(100000)
}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

//вставка
//func BenchmarkSelectionSort(b *testing.B) {
//	b.Run("small arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			ar := generateSlice(10, 10)
//			b.StartTimer()
//			selectionSort(ar)
//			b.StopTimer()
//		}
//	})
//
//	b.Run("middle arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			ar := generateSlice(100, 1000)
//			b.StartTimer()
//			selectionSort(ar)
//			b.StopTimer()
//		}
//	})
//
//	b.Run("big arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			ar := generateSlice(10000, 100000)
//			b.StartTimer()
//			selectionSort(ar)
//			b.StopTimer()
//		}
//	})
//}
//вставка

//
func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
