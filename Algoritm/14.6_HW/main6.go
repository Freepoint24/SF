package main

import "fmt"


func main() {

	var n1 int = 0       //размер первого массива
	var n2 int = 0       //размер второго массива
	var finalarray []int //итоговый массив
	var err error
	var arr1 map[string]int


	type void struct{}
	var arr1 void
	//firstarray := map[int]void{1: member, 9: member, 12: member, 16: member, 13: member, 11: member, 14: member, 18: member, 8: member}
	arr1 := map[int]void{1: member, 9: member, 12: member, 16: member, 13: member, 11: member, 14: member, 18: member, 8: member}

	//secondarray := map[int]void{9: member, 7: member, 15: member, 8: member, 14: member, 19: member, 34: member, 18: member, 13: member, 11: member}
	arr2 := map[int]void{9: member, 7: member, 15: member, 8: member, 14: member, 19: member, 34: member, 18: member, 13: member, 11: member}




	//readInt("Enter first array size: \n", &n1) //ввести размер первого массива
	fmt.Println("Enter first array size:")
	_, err = fmt.Scanln(&n1)
	arr1 := make([]int, n1)
	if err != nil {
		fmt.Printf("размер массива 1 должен быть положительним числом \n")
		return
	}

	//readInt("Enter second array size: \n", &n2) //ввести рамзер второго массива
	fmt.Println("Enter second array size:")
	//fmt.Scan(&n2)
	_, err = fmt.Scanln(&n2)
	arr2 := make([]int, n2)
	if err != nil {
		fmt.Printf("размер массива 2 должен быть положительним числом \n")
		return
	}


	fmt.Printf("Enter first array: \n") //заполняем первый массив
	var k1 int
	for i := 0; i < n1; i++ {
		fmt.Scan(&k1)
		arr1[i] = k1
	}


	fmt.Printf("Enter first array: \n") //заполняем второй массив
	var k2 int
	for i := 0; i < n2; i++ {
		fmt.Scan(&k2)
		arr2[i] = k2
	}


	//ищем совпадения в двух массивах
	for _, a1 := range arr1 {
		for _, a2 := range arr2 {
			if a1 == a2 {
				finalarray = append(finalarray, a2)
			}
		}
	}
	//fmt.Print(arr1, "\n")
	//fmt.Print(arr2, "\n")

	fmt.Println(finalarray) //итоговый массив

}

main_14.6_HW.go