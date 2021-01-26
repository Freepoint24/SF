package main

import (
	"container/list"
	"fmt"
)

func main() {
	//запись шахматных ходов на основе связного списка
	//создаем пустой список
	queue := list.New()

	queue.PushFront("e4, e5") // Добавление в очередь каждого хода
	queue.PushBack("c4, c6")
	queue.PushBack("h5, f6")
	queue.PushBack("xf7")

	//Вывод первого и последнего хода
	x := queue.Front() // Первый ход
	fmt.Println("Первый ход:", x.Value)
	y := queue.Back() // Последний ход
	fmt.Println("Последний ход:", y.Value)
	//z := queue.Len()
	//всего ходов
	//fmt.Println(*queue)
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}

	// Печать всех ходов

}
