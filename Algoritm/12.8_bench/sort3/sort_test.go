package main

import (
	"math/rand"
	"testing"
)

//запускать командой go test -bench=. ./...

//12.3.2
//алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
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
	//Добавить проверки для других размеров входных данных, отличающихся от примера, попробовать использовать
	//большие и меньшие максимальные числа.
	//малый размер данных
	//	b.Run("SMOLL_small arrays", func(b *testing.B) {
	//		b.StopTimer()
	//		for i := 0; i < b.N; i++ {
	//			ar := generateSlice(10, 100)
	//			b.StartTimer()
	//			bubbleSort(ar)
	//			b.StopTimer()
	//		}
	//	})
	//большой размер данныйх
	//правал теста
	//	b.Run("BIG_big arrays", func(b *testing.B) {
	//		b.StopTimer()
	//		for i := 0; i < b.N; i++ {
	//			ar := generateSlice(100000, 400000)
	//			b.StartTimer()
	//			bubbleSort(ar)
	//			b.StopTimer()
	//		}
	//	})
}

////////////////////////////////////////////////////////////////////////////////////////
//Задание 12.4.1
//Реализуйте сортировку выбором Selection sort, работающую «слева направо»
func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs()
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
	//Добавить проверки для других размеров входных данных, отличающихся от примера, попробовать использовать
	//большие и меньшие максимальные числа.
	//малый размер данных
	//b.Run("SMOLL_small arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(5, 10)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
	//большой размер данныйх
	//провал теста
	//b.Run("BIG_big arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(100000, 400000)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
//Задание 12.5.1
//Insertion sort (Сортировка выбором)
func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	// провал теста
	//fmt.Println("BenchmarkInsertionSort/ - small arrays - Test error")
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
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	//Добавить проверки для других размеров входных данных, отличающихся от примера, попробовать использовать
	//большие и меньшие максимальные числа.
	//малый размер данных
	//b.Run("SMOLL_small arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(5, 10)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
	//большой размер данныйх
	//провал теста
	//b.Run("BIG_big arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(100000, 400000)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
}

////////////////////////////////////////////////////////////////////////////////////////////////
//ЗАДАНИЕ 12.6.1
//Реализуйте сортировку слиянием, Merge sort
func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		//fmt.Println("BenchmarkMergeSort/ - small arrays - Test error")
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	//провал теста
	b.Run("middle arrays", func(b *testing.B) {
		//fmt.Println("BenchmarkMergeSort/ - middle arrays - Test error")
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
	//Добавить проверки для других размеров входных данных, отличающихся от примера, попробовать использовать
	//большие и меньшие максимальные числа.
	//малый размер данных
	//b.Run("SMOLL_small arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(5, 10)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
	//большой размер данныйх
	//провал теста
	//b.Run("BIG_big arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(100000, 400000)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
//Задание 12.7.1
//Реализуйте быструю сортировку.
func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs()
	//провал теста
	b.Run("small arrays", func(b *testing.B) {
		//fmt.Println("BenchmarkQuickSort/ - small arrays - Test error")
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
	//Добавить проверки для других размеров входных данных, отличающихся от примера, попробовать использовать
	//большие и меньшие максимальные числа.
	//малый размер данных
	//b.Run("SMOLL_small arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(5, 10)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
	//большой размер данныйх
	//провал теста
	//b.Run("BIG_big arrays", func(b *testing.B) {
	//	b.StopTimer()
	//	for i := 0; i < b.N; i++ {
	//		ar := generateSlice(100000, 400000)
	//		b.StartTimer()
	//		bubbleSort(ar)
	//		b.StopTimer()
	//	}
	//})
}

/////////////////////////////////////////////////////////////////////////////////////
func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
