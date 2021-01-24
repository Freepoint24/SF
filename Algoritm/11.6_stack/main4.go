package main

import "fmt"

func main() {
	//журнал шахматых ходов
	var allMove []string

	allMove = append(allMove, "A2") // Push, добавляем ходы
	allMove = append(allMove, "M5")
	allMove = append(allMove, "J4")
	allMove = append(allMove, "C8")
	allMove = append(allMove, "F1")
	allMove = append(allMove, "D5")
	// вывести все записанные ходы
	fmt.Println("Все ходы", allMove)
	// Pop, выводит последний ход
	n := len(allMove) - 1
	fmt.Println("Последний ход:", allMove[n])
}
