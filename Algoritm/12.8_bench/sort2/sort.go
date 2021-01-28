package main

//запускать командой go test -bench=. ./...
func bubbleSort(ar []int) {
	// sort on the right - по убыванию
	for i := 0; i < (len(ar) - 1); i++ {
		for j := 0; j < ((len(ar) - 1) - i); j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}

//
//func selectionSort(ar []int) {
//		for i := 0; i < len(ar); i++ {
//		minIndex := i
//		j := i + 1
//		for j < len(ar) {
//			if ar[j] < ar[minIndex] {
//				minIndex = j
//			}
//			j = j + 1
//		}
//		ar[i], ar[minIndex] = ar[minIndex], ar[i]
//		}
//
//}
