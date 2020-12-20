package main2

import "fmt"

func main() {
	var x = [9]int{10, 11, 12, 13, 14, 15, 16, 17, 18}

	fmt.Println(x[3])

	example(x)
	fmt.Println(x[3])
}

func example(arr [9]int)  {
	arr[3] = 4

}


