package main

import (
	"fmt"
	"log"
)

func main() {

	//Подготовьте кусочек струн.
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	//Создайте дерево и заполните его значениями.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	//Распечатайте отсортированные значения.
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	//Найдите ценности.
	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	//Удалить значение.
	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	//Особый случай: одноузловое дерево. (Посмотрите, Tree.Deleteпочему это особый случай.)
	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}
