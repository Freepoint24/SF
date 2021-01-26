package main

import "fmt"

//написать код, который динамически расчитывает среднюю температуру от 10 метеостанций
//Данные от каждой станции постоянно обновляются.
//изменения температуры от какой-либо метеостанции — это ввод соответствующей температуры через консоль
//и номера метеостанции, от которой она получена.
//При каждом изменении температуры на любой станции ваша программа должна сразу вывести среднюю температуру в городе.

func main() {
	// обяъявляем массив размерностью 10, по количеству метеостанций
	var dataTemperatura [10]int
	var err error
	var nomerStan int
	var temperatura int
	for {
		//ввод номера станции
		fmt.Print("Введите номер станции от 1 до 10: ")
		_, _ = fmt.Scanln(&nomerStan)
		switch nomerStan {
		case 1:
			nomerStan = 0
		case 2:
			nomerStan = 1
		case 3:
			nomerStan = 2
		case 4:
			nomerStan = 3
		case 5:
			nomerStan = 4
		case 6:
			nomerStan = 5
		case 7:
			nomerStan = 6
		case 8:
			nomerStan = 7
		case 9:
			nomerStan = 8
		case 10:
			nomerStan = 9
		default:
			fmt.Printf("Ошибка ввода \n")
			_, _ = fmt.Scanln(&nomerStan)
		}

		//ввод температуры станции
		fmt.Print("Введите температуру станции: ")
		_, err = fmt.Scanln(&temperatura)
		if err != nil {
			fmt.Printf("ошибка ввода данных")
			return

		}
		//определяем от какой станции вводим температуру
		dataTemperatura[nomerStan] = temperatura
		//суммируем все данные
		total := 0
		for _, val := range dataTemperatura {
			total += val
		}
		//вычисляем среднюю температуру и выводим на экран
		fmt.Println("Среднее значение: ", total/10)
		// вывод всех текущих данных
		//fmt.Println(dataTemperatura)
	}
}
