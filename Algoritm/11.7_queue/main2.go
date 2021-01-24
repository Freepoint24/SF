package main

import (
	"container/list"
	"fmt"
)

func main() {
	//запись шахамтных ходов на основе связного списка
	//создаем пустой список
	queue := list.New()

	queue.PushBack("e4, e5") // Добавление в очередь каждого хода
	queue.PushBack("c4, c6")
	queue.PushBack("h5, f6")
	queue.PushBack("xf7")

	//Вывод первого и последнего хода
	x := queue.Front() // Первый ход
	fmt.Println("Первый ход:", x.Value)
	y := queue.Back() // Послдений ход
	fmt.Println("Последний ход:", y.Value)

}
