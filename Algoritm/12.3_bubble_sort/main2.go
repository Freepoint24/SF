package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	BubbleSort(ar)
	fmt.Println(ar)
}

func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, i, j)
			}
		}
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
