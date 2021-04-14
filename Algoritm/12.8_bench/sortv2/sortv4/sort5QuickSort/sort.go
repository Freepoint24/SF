package sort

import "math/rand"

//5////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 12.7. Быстрая сортировка quickSort
Сложность
В лучшем случае: O(n log(n)).
В среднем случае: O(n log(n)).
В худшем случае: O(n2).
Емкостная, в худшем: O(1).

*/

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
