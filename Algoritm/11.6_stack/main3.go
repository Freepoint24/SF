package main

import (
	"fmt"
)

// объявить слайс
// добоавит три элемета
// вернуть последний

func main() {
	//записывам первые сто шагов шахматной партии
	//push
	allMove := [100]string{}
	allMove[0] = "A5"
	allMove[1] = "C3"
	allMove[2] = "F1"
	allMove[3] = "J4"
	// выводи все записанные шаги
	//fmt.Println(allMove)
	//выводим последний шаг
	//pop

	for len(allMove) > 0 {
		s := len(allMove) - 1
		fmt.Print(allMove[s])
		//allMove = allMove[:s]
	}

}
