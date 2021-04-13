package sort

//1///////////////////////////////////////////////////////////////////////////////////////////////
/*12.3.2 алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
/Сложность bubbleSort
/В лучшем случае: O(n).
/В среднем случае: O(n2).
/В худшем случае: O(n2) Е.
Ёмкостная, в худшем: O(1).
*/
func bubbleSort(s1 []int) {
	length := len(s1)                   //переменная length - длина слайса s1
	for i := 0; i < (length - 1); i++ { // переменная i - количество проходов по слайсу
		for j := 0; j < ((length - 1) - i); j++ { //переменая j - индекс элемента слайса
			if s1[j] > s1[j+1] { //если предыдущее значение меньше следующего, то
				s1[j], s1[j+1] = s1[j+1], s1[j] //поменять их местами
			}
		}
	}

}

//2////////////////////////////////////////////////////////////////////////////////////////////////
/*Задание 12.4.1 сортировку выбором Selection sort, работающую «слева направо»
Сложность
В лучшем случае: O(n2)
В среднем случае: O(n2)
В худшем случае: O(n2)
Ёмкостная, в худшем: O(1)
*/
func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ { //переменная i - количество прогонов по слайсу
		minIndex := i     //переменная minIndex - минимальное значение в перебираемом слайсе
		j := i + 1        //индекс начала перебираемого слайса
		for j < len(ar) { //для j меньше длины слайса,
			if ar[j] < ar[minIndex] { //сравнвимваем значение в индексе j c минимальным значением
				minIndex = j //находим минимальное значение в слайсе
			}
			j = j + 1 //сравниваем значение хранящееся в следующем индексе слайса
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i] //меняем местами минимальное значение и значение в текущем индекс слайса
	}
}

//3//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*Задание 12.5.1 Сортировка вставкой InsertionSort
Сложность
В лучшем случае: O(n).
В среднем случае: O(n2).
В худшем случае: O(n2).
Емкостная, в худшем: O(1).
*/

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ { //переменная i - количество проходов по слайсу
		x := ar[i]                         //переменная x - i-тый индекс массива
		j := i                             //переменная j присвоить значение i
		for ; j >= 1 && ar[j-1] > x; j-- { //если j больше или равно 1 и предыдущий элемент больше итого значения
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}

//4////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 12.6. Сортировка слиянием mergeSort
Сложность
В лучшем случае: O(n log(n)).
В среднем случае: O(n log(n)).
В худшем случае: O(n log(n)).
Емкостная, в худшем: O(n).
*/

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

//5////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 12.7. Быстрая сортировка quickSort
Сложность
В лучшем случае: O(n log(n)).
В среднем случае: O(n log(n)).
В худшем случае: O(n2).
Емкостная, в худшем: O(1).

*/

func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)
	quickSort(ar[:split])
	quickSort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]
	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
