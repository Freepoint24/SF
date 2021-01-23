package main

import "fmt"

func main() {
	for _, r := range "слово" {
		fmt.Println(r)         //юникод номер в таблице юникод
		fmt.Println(string(r)) //
	}

}
