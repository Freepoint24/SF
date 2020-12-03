package main

import "fmt"

// структура начало
type A struct {
	i int // поле i тип int (экземляр структуры)
}

// структура конец
//инициализация структуры методом square тип int
// приемник (переменная) называется a принадлежит структуре А
func (a *A) increase() { // *A это указатель на структуру
	a.i = a.i + 15
}

func main() {
	a := A{i: 10}
	a1 := A{i: 5}

	a.increase()
	fmt.Println(a.i)

	a.increase()
	fmt.Println(a1.i)
}
