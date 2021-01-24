package main

import "fmt"

func greet(name string) {
	fmt.Printf("hello, %v !\n", name)
	bye := greet2(name) // функция greet2 вернула функцию
	fmt.Printf("getting ready to say bye...\n")
	bye() // тут мы вызвали эту функцию
}

func greet2(name string) func() {
	fmt.Printf("how are you, %v ?\n", name)
	bye := func() { // тут мы создали анонимную функцию и присвоили её переменной
		fmt.Println("ok bye!")
	}
	return bye // тут мы просто возвращаем переменную
}

func main() {
	greet("Sasha")
}
