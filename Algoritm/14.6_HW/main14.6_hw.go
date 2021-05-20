package main

import (
	"fmt"
)


func main() {

	var n1 int = 0       //размер первого массива
	var n2 int = 0       //размер второго массива
	var finalarray []int //итоговый массив
	var err error

// ввод размера первого массива
	fmt.Println("Enter first array size:")
	_, err = fmt.Scanln(&n1)
	if err != nil {
		fmt.Printf("Размер массива 1 должен быть положительным числом \n")
		return
	}

// ввод размера второго массива
	fmt.Println("Enter second array size:")
	_, err = fmt.Scanln(&n2)
	if err != nil {
		fmt.Printf("Размер массива 2 должен быть положительным числом \n")
		return
	}

// ввод данных первого массива
	fmt.Println("Enter first array:")
	mapp1 := make(map[int]string)
			for i := 0; i < n1; i++ {
			var num int
			var text string

		fmt.Scanf("%d %s", &num, &text)
		mapp1[num] = text
	}

	// fmt.Print(mapp1, "\n") //////первый массив


// ввод данных второго массива
	fmt.Println("Enter second array:")
	mapp2 := make(map[int]string)
	for i := 0; i < n2; i++ {
		var num int
		var text string

		fmt.Scanf("%d %s", &num, &text)
		mapp2[num] = text
	}

	//fmt.Print(mapp2, "\n") ///// второй массив

//ищем совпадения в двух массивах
		for k := range mapp1 {
		if _, ok := mapp2[k]; ok {
			finalarray = append(finalarray, k)
		}
	}

// итоговый массив
	fmt.Println(finalarray)
}




