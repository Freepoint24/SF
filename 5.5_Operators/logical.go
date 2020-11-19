package main

import "fmt"

func main()  {

	t := true
	f := false

	// Логическое Умножение или логическое "И". Выведет false (0 * 1 = 0)
	fmt.Println(t && f)
	// Логическое Сложение иди длшическое "ИЛИ". Вывежеь true (1+ 0 = 1)
	fmt.Println(t || f)
	// Логическое отрицание или "НЕ". Выведет значени обратное имеющемуся. В данном случае - false.
	// !true = false
	fmt.Println(!t)
}