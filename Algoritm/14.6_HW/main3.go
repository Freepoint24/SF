package main

import "fmt"

func main() {

	var n1 int = 0       //размер первого массива
	var n2 int = 0       //размер второго массива
	var finalarray []int //итоговый массив

	readInt("Enter first array size: \n", &n1) //ввести размер первого массива
	arr := make([]int, n1)

	readInt("Enter second array size: \n", &n2) //ввести рамзер второго массива
	arr2 := make([]int, n2)

	fmt.Printf("Enter first array: \n") //заполняем первый массив
	for i := 0; i < n1; i++ {
		var k1 int = 0
		readInt("", &k1)
		arr[i] = k1
	}

	fmt.Printf("Enter second array: \n") //заполнить второй массив
	for i := 0; i < n2; i++ {
		var k2 int = 0
		readInt("", &k2)
		arr2[i] = k2
	}

	for _, a1 := range arr {
		for _, a2 := range arr2 {
			if a1 == a2 {
				finalarray = append(finalarray, a2)
			}
		}
	}
	fmt.Println(finalarray)
}

func readInt(s string, x *int) {
	fmt.Print(s + "")
	var y int = 0
	fmt.Scan(&y)
	*x = y
}
