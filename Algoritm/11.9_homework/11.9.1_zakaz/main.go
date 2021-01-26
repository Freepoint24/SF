package main

import (
	"container/list"
	"fmt"
)

//Поступление нового заказа будет увеличивать счетчик заказов на +1
//Значение счетчика передается в конец очереди - тип данных динамическая очередь на основе связного списка
//на обработку заказов значине забирается из начала очереди в функцию выполнения заказов
//если покупатель передумал, то удаляется соотвтстсвующий номер заказа из очереди

func main() {
	//записать заказы интернет магазина
	//создаем пустой список
	queue := list.New()

	queue.PushFront(1) // Добавление нового заказа
	queue.PushBack(5)
	queue.PushBack(6)
	queue.PushBack(7)

	//Вывод первого заказа для обработки
	x := queue.Front() // первый заказ для обработки
	fmt.Println("Первый заказ для обработки:", x.Value)
	//y := queue.Back() // Последний поступивший заказ
	//fmt.Println("Последний ход:", y.Value)

	//все заказы
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}

}
