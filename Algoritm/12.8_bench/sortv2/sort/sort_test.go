package sort

import (
	"math/rand"
	"testing"
)

//go test -bench=. -benchmem > new.txt записать в файл
//запускать командой go test -bench=. ./...
//go test -bench=. -benchmem  покажет алокацию памяти
//1/////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
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
}

//2//////////////////////////////////////////////////////////////////////////////////////
func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.StopTimer()    //
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

//3////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.StopTimer()    //
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

//4//////////////////////////////////////////////////////////////////////////////////////////////////////////

func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.StopTimer()    //
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
}

//5/////////////////////////////////////////////////////////////////////////////////////////////////////////

func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.StopTimer()    //
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}

///////////////////////////////////////////////////////
func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max //генерация случайного числа
	}
	return ar
}
