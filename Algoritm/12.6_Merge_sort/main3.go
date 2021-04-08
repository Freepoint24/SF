package main

func mergSort3(arr []int []int)  {
	item := make([]int, len(arr))
	copy(items, arr)
	doMergeSort(items)
	return items
	}
func doMergSort(items []int)  {
	lenght := len(items)
	if lenght == 1 {
		return
	}

	ILeft := lenght / 2
	left := make([]int, ILeft)
	copy(left, items[:ILeft])

	IRight := lenght / 2
	left := make([]int, IRight)
	copy(left, items[:ILeft])

	doMergSort(left)
	doMergSort(right)

	merge(left, right, items)

}

func merge3(left []int, right []int, items []int)  {
	l := 0
	r := 0
	i := 0

	for (l < len(left)) && (r<len(right)){
		if left[l] < right[r] {
			result[i] = left[i]
			l++
		} else {
			result[i] = right[r]
				r++
		}
		i++
	}
	lenght := len(left) - l
	copy(result[i:i + lenght], left[l:])
	i = i + lenght
	lenght = len(lenght, right[r:])

}


func main (){
	items := []int{4, 1, 5, }
	sortitem := mergSort(items)
}

