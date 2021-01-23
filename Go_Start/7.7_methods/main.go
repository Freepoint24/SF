package main

import "fmt"

// структура начало
type A struct {
	i int // поле i тип int (экземляр структуры)
}

// структура конец
//инициализация структуры методом square тип int
// приемник (переменная) называется a принадлежит структуре А
func (a A) square() int {
	return a.i * a.i
}

func main() {
	a := A{i: 10}
	a1 := A{i: 5}

	fmt.Println(a.square())
	fmt.Println(a1.square())
}
