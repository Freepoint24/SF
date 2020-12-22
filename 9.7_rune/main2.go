package main

import "fmt"

func main() {
	str := "слово"
	runes := []rune(str)
	runes2 := []rune("слово")

	fmt.Println(len(runes))
	fmt.Println(string(runes2[1]))
	fmt.Println(len(str))

}
