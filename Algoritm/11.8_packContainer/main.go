package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем кольцо из пяти элементов
	r := ring.New(5)
	// Инициализируем элементы кольца значениями
	// Обратите внимание мы обошли кольцо полностью один раз!
	// Полностью!
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	// Теперь кто-то из вас может подумать, что повторная
	// итерация по
	// элементам приведет к ошибке
	// ведь мы уже в конце кольца, куда нам дальше-то двигаться?

	// Из конца списка мы вновь возвращаемся автоматически в
	// начало и
	// перемещаемся вперед по кольцу на два элемента
	for j := 0; j < 2; j++ {
		r = r.Next()
	}
	// А теперь внимание!
	// Мы создаем новое кольцо - просто указав
	// ссылку на следующий элемент уже существующего кольца
	newRing := r.Next()
	r = r.Next()
	fmt.Printf("%d %d", newRing.Value, r.Value)
}