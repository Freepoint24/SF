package main

import (
	"math/rand"
	"testing"
)

//запускать командой go test -bench=. ./...
//go test -bench=. -benchmem  покажет аалокацию

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs() //аллокация памяти
	b.StopTimer()    //
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
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
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
