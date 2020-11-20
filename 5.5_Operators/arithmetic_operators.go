package main

import "fmt"

func main() {

	// Оператор сложения
	sum := 2 + 3

	// Оператор вычитания
	subtraction := 3 - 2
	// Оператор умножения
	multiplication := 2 * 2

	// Оператор целоечисленного деления. На 0 нельзя делить даже в программировании, иначе будет ошибка исполнения программы.
	// Резелтатом данного деления будет целое число.
	division := 2 / 2

	// Чтобы получить остаток, следует использовать оператор взятия остатка от деления
	divisionRemainder := 3 % 2

	// Оператор инкремента. Увеличивает значение функции на единицу
	increment := 2
	increment++

	// Оператор декремента. Уменьшает значение функции на единицу
	decrement := 3
	decrement--

	fmt.Println(sum, subtraction, multiplication, division, divisionRemainder, increment, decrement)
}
