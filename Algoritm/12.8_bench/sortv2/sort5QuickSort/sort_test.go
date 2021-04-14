package sort

import (
	"math/rand"
	"testing"
)

////go test -bench=. -benchmem > new1.txt записать в файл
////запускать командой go test -bench=. ./...
////go test -bench=. -benchmem  покажет алокацию памяти

//5/////////////////////////////////////////////////////////////////////////////////////////////////////////

func BenchmarkQuickSort(b *testing.B) {
	//b.ReportAllocs() //аллокация памяти
	b.StopTimer() //
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ { //b.N спецаильная переменная, от 0 до b.N выполнять цикл
			ar := generateSlice(10, 100)
			b.StartTimer()
			b.ReportAllocs()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		//b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 10)
			b.StartTimer()
			b.ReportAllocs()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		//b.ReportAllocs()
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			b.ReportAllocs()
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
