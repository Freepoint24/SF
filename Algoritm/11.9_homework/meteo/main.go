package main

import "fmt"

//написать код, который динамически рассчитывает среднюю температуру от 10 метеостанций
//Данные от каждой станции постоянно обновляются.
//изменения температуры от какой-либо метеостанции — это просто ввод соответствующей температуры через консоль
//и номера метеостанции, от которой она получена.
//При каждом изменении температуры на любой станции ваша программа должна сразу вывести среднюю температуру в городе.

//массив 10 штук
//серднюю в массиве найти

func main() {
	// обяъявляем массив размерностью 10, по количество метеостаций
	var x [10]int
	var err error
	//ввод температуры стацния №1
	fmt.Print("Введите температуру 1 станции: ")
	_, err = fmt.Scanln(&x[0])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}

	//ввод температуры стацния №2
	fmt.Print("Введите температуру 2 станции: ")
	_, err = fmt.Scanln(&x[1])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}

	//ввод температуры стацния №3
	fmt.Print("Введите температуру 3 станции: ")
	_, err = fmt.Scanln(&x[2])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №4
	fmt.Print("Введите температуру 4 станции: ")
	_, err = fmt.Scanln(&x[3])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №5
	fmt.Print("Введите температуру 5 станции: ")
	_, err = fmt.Scanln(&x[4])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №6
	fmt.Print("Введите температуру 6 станции: ")
	_, err = fmt.Scanln(&x[5])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №7
	fmt.Print("Введите температуру 7 станции: ")
	_, err = fmt.Scanln(&x[6])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №8
	fmt.Print("Введите температуру 8 станции: ")
	_, err = fmt.Scanln(&x[7])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №9
	fmt.Print("Введите температуру 9 станции: ")
	_, err = fmt.Scanln(&x[8])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}
	//ввод температуры стацния №10
	fmt.Print("Введите температуру 10 станции: ")
	_, err = fmt.Scanln(&x[9])
	if err != nil {
		fmt.Printf("ошибка ввода данных")
		return
	}

	total := 0
	for _, val := range x {
		total += val
	}
	//вычисляем средню температуру и выводим на экран
	fmt.Println("Среднее значение: ", total/10)
	// вывод всех данных
	fmt.Println(x)
}
