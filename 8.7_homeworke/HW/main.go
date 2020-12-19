package main

import (
	"fmt"
	"home/electronic"
)

func main() {
	//f := func(b float64) float64 {return 0.0}
	//if f(5) > 0 {
	//	fmt.Println(f(5))
	//}

	nIphone := electronic.NapplePhone("X11")
	printCharacteristics(nIphone)

	nAndroid := electronic.NandroidPhone("Xiaomi", "A6")
	printCharacteristics(nAndroid)

	nRadio := electronic.NradioPhone("BBK", "18", 9)
	printCharacteristics(nRadio)
}

func printCharacteristics(phone electronic.Phone) {
	//fmt.Printf("Brand: %v Model: %v Type: %v \n", phone.Brand(), phone.Model(), phone.Type())
	fmt.Println("Brand: ", phone.Brand())
	fmt.Println("Model:", phone.Model())
	fmt.Println("Type:", phone.Type())
	//fmt.Println("")

	//switch phone := phone.(type) {
	//	case electronic.Smartphone:
	//	fmt.Println("os", phone.OS())
	//}
	smart, ok := phone.(electronic.Smartphone)
	if ok {
		fmt.Printf("os %s \n", smart.OS())
	}
	radio, ok := phone.(electronic.StationPhone)
	if ok {
		fmt.Printf("buttons %v \n", radio.ButtonsCount())
	}

	fmt.Println()
}
