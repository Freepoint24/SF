package main

import "fmt"

func main() {

	const maxCounter = 100

	var counter int

	for working := true; working != false; {
		counter++

		if counter >= 100 {
			fmt.Println(counter)

			working = false
		}
	}

}
