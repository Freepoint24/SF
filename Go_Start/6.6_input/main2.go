package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// пытаемся сконвертировать строку, содержащую число -42 в тип int
	i, err := strconv.Atoi("-42")
	fmt.Println(i, err)
	// конвертируем число - 42 в строку
	s := strconv.Itoa(-42)
	fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	text, _ := reader.ReadString('\n')
	fmt.Println("Введенная строка:", text)
}
