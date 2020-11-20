package test

import "fmt"

// Функция названа с заклавной буквы, ее можно будет вызвать из других пакетов
func Public()  {
	fmt.Println("I`m public")
}

// Функция названа со строчной буквы, ее можно будет вызывать только в текущем пакете
func private()  {
	fmt.Println("I`m private")
}
