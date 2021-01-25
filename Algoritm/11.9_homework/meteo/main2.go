package main

import (
	"fmt"
	//"strconv"
)

//написать код, который динамически рассчитывает среднюю температуру от 10 метеостанций
//Данные от каждой станции постоянно обновляются.
//изменения температуры от какой-либо метеостанции — это просто ввод соответствующей температуры через консоль
//и номера метеостанции, от которой она получена.
//При каждом изменении температуры на любой станции ваша программа должна сразу вывести среднюю температуру в городе.

//массив 10 штук
//серднюю в массиве найти

func main() {
	// обяъявляем массив размерностью 10, по количество метеостанций
	var x [10]int
	var err error
	var nomerStan int
	var temperatura int
	//var y int

	//ввод номер стацния
	fmt.Print("Введите номер станции: ")
	_, _ = fmt.Scanln(&nomerStan)
	switch nomerStan {
	case 1:
		nomerStan = 0 //"x[0]"
	case 2:
		nomerStan = 1 //"x[1]"
	case 3:
		nomerStan = 2 //"x[2]"
	case 4:
		nomerStan = 3 //"x[3]"
	default:
		fmt.Printf("Ошибка ввода \n")
		return
	}
	//fmt.Println(nomerStan)

	//ввод температуры стацния
	fmt.Print("Введите температуру станции: ")
	_, err = fmt.Scanln(&temperatura)
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//y := strconv.Atoi(nomerStan)

	//y, err = strconv.Atoi(nomerStan)
	//if err != nil {
	//	fmt.Printf("некорректное число")
	//}

	x[nomerStan] = temperatura

	total := 0
	for _, val := range x {
		total += val
	}
	//вычисляем средню температуру и выводим на экран
	fmt.Println("Среднее значение: ", total/10)
	// вывод всех данных
	fmt.Println(x)
}
