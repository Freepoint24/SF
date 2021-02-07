package main

import (
	"errors"
)
//Бинарное дерево
type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func main() {

	//вставка элемента,
	//Insert вставляет новые данные в дерево в позиции, определяемой значением поиска.
	//Возвращаемые значения:
	//true если данные были успешно вставлены
	//false если значение данных уже существует в дереве.
	func(n *Node) Insert(value,	data string) error{
		if n == nil{
		return errors.New("Cannot insert a value into a nil tree")
	}

		switch{   		//	Если данные уже есть в дереве, вернитесь.
	case value == n.Value:
		return nil 		//	Если значение данных меньше, чем значение текущего узла, и если левый дочерний узел меньше nil, вставьте новый левый дочерний узел. Другой вызов Insertв левом поддереве.
	case value < n.Value:
		if n.Left == nil{
		n.Left = &Node{Value: value, Data: data}
		return nil
	}
		return n.Left.Insert(value, data) 		//	Если значение данных больше, чем значение текущего узла, сделайте то же самое, но для правого поддерева.
	case value > n.Value:
		if n.Right == nil{
		n.Right = &Node{Value: value, Data: data}
		return nil
	}
		return n.Right.Insert(value, data)
	}
		return nil
	}


}