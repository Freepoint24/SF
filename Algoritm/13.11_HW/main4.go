package main

//реализовать 1.
//Структура: двоичное дерево.
//Функции: вставка элемента, удаление элемента, поиск элемента, печать элементов дерева.

import (
	//"errors"
	"errors"
	"fmt"
)

//TreeNode банирное дерево структура
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
	Data  string
}

func main() {

	t := &TreeNode{val: 66} //добавить элемент в дерево

	//заполняем дерево начальными значениями
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	//t.Find(12) //Поиск значения в дереве
	//t.Delete(0) //удаление значения из дерева

	//Печать элементов дерева
	t.PrintTreeNode()
	fmt.Println("")
}

//распечатать элементы по порядку
func (t *TreeNode) PrintTreeNode() {
	if t == nil {
		return
	}
	t.left.PrintTreeNode()
	fmt.Print(t.val)
	t.right.PrintTreeNode()
}

func (t *TreeNode) Insert(val, data string) error {

	if t == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	// If the data is already in the tree, return.
	case val == t.val:
		return nil
	// If the data value is less than the current node's value, and if the left child node is `nil`, insert a new left child node. Else call `Insert` on the left subtree.
	case val < t.val:
		if t.left == nil {
			t.left = &TreeNode{val: val, Data: data}
			return nil
		}
		return t.left.Insert(val, data)
	// If the data value is greater than the current node's value, do the same but for the right subtree.
	case val > t.val:
		if t.right == nil {
			t.right = &TreeNode{val: val, Data: data}
			return nil
		}
		return t.right.Insert(val, data)
	}
	return nil
}

////Insert вставить новый элемент
//func (t *TreeNode) Insert(value int) error {
//	//if t == nil {
//	//	return errors.New("Tree is nil")
//	//}
//	//if t.val == value {
//	//	return errors.New("This node value already exists")
//	//}
//	if t.val > value {
//		if t.left == nil {
//			t.left = &TreeNode{val: value}
//			return nil
//		}
//		return t.left.Insert(value)
//	}
//	if t.val < value {
//		if t.right == nil {
//			t.right = &TreeNode{val: value}
//			return nil
//		}
//		return t.right.Insert(value)
//	}
//	return nil
//}
//
////Delete удаление элемента из двоичного дереве
//func (t *TreeNode) Delete(value int) {
//	t.remove(value)
//}
//
//func (t *TreeNode) remove(value int) *TreeNode {
//	if t == nil {
//		return nil
//	}
//	if value < t.val {
//		t.left = t.left.remove(value)
//		return t
//	}
//	if value > t.val {
//		t.right = t.right.remove(value)
//		return t
//	}
//	if t.left == nil && t.right == nil {
//		t = nil
//		return nil
//	}
//	if t.left == nil {
//		t = t.right
//		return t
//	}
//	if t.right == nil {
//		t = t.left
//		return t
//	}
//
//	smallestValOnRight := t.right
//	for {
//		//найти наименьшее значение с правой стороны
//		if smallestValOnRight != nil && smallestValOnRight.left != nil {
//			smallestValOnRight = smallestValOnRight.left
//		} else {
//			break
//		}
//	}
//
//	t.val = smallestValOnRight.val
//	t.right = t.right.remove(t.val)
//	return t
//}
//
////Find поиск элемента в бинарном дереве
//func (t *TreeNode) Find(value int) (TreeNode, bool) {
//
//	if t == nil {
//		return TreeNode{}, false
//	}
//		switch {
//	case value == t.val:
//		return *t, true
//
//	case value < t.val:
//		return t.left.Find(value)
//	default:
//		return t.right.Find(value)
//
//	}
//
//}
