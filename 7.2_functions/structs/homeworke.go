package main

type pet struct {
	name string
	age  int
}

type cat struct {
	pet
	woolColor string //цвет шерсти
}

func NewCat(name string, age int, woolColor string) cat {
	return cat{pet: pet{name: name, age: age}, woolColor: woolColor}
}
