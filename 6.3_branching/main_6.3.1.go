package main

import "fmt"

	const passportAge = 23
	var age int
	//var mustHavePassport bool
func main() {
	if age >= 14 {
	//mustHavePassport = true
		fmt.Println(true)
	}else if age < 14 {
	fmt.Println(false)
		}
	}
