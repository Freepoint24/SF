package main

import "fmt"

func main() {
	producePanic()

	defer func() {
		panicMsg := recover()
		if panicMsg != "" {
			fmt.Printf("after panic with msg: %s", panicMsg)
		}
	}()
}

func producePanic() {
	fmt.Println("before panic")
	panic("critical error")
}
