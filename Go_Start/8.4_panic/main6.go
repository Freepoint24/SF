package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("panic occured: %s", r)
		} else {
			fmt.Println("no panic!")
		}
	}()
	var i *int

	producePanic(i)
}

func producePanic(i *int) {
	fmt.Println(*i)
}
