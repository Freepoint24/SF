package main

import "fmt"

func main()  {
	i := 3
	if i > 0 {
		fmt.Println(0)
	}else if i > 1 {
		fmt.Println(1)
	}else if i > 1 {
		fmt.Println(2)
	}else if i > 1 {
		fmt.Println(3)
	}

	switch i {
	case 0: fmt.Println(0)
	 fallthrough
	case 1: fmt.Println(1)
	case 2: fmt.Println(2)
	case 3: fmt.Println(3)
	default:
		fmt.Println("default case")
		fmt.Println(i)
	}
}