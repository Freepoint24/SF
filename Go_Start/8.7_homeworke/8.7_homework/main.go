package main

import (
	"fmt"
	"home/electronic"
)

//создаеем по одному экземпляру каждой структуры из пакета electronic.
// функция которая принимает на вход параметр типа Phone (интерфейс).
//Функция printCharacteristics выводит бренд, модель и тип телефона.
func main() {

	nIphone := electronic.NapplePhone("X11")
	printCharacteristics(nIphone)

	nAndroid := electronic.NandroidPhone("Xiaomi", "A6")
	printCharacteristics(nAndroid)

	nRadio := electronic.NradioPhone("BBK", "18", 9)
	printCharacteristics(nRadio)
}

//вызываем функцию printCharacteristics для каждого созданного экземпляра структур applePhone, androidPhone и radioPhone.
func printCharacteristics(phone electronic.Phone) {
	fmt.Println("Brand: ", phone.Brand())
	fmt.Println("Model:", phone.Model())
	fmt.Println("Type:", phone.Type())

	//если телефон также является стационарным телефоном, то должно быть выведено количество его кнопок,
	radioPhone, ok := phone.(electronic.StationPhone)
	if ok {
		fmt.Printf("buttons %v \n", radioPhone.ButtonsCount())
	}
	//если телефон является смартфоном, то должно быть выведено название его операционной системы.
	smartPhone, ok := phone.(electronic.Smartphone)
	if ok {
		fmt.Printf("os %s \n", smartPhone.OS())
	}
	fmt.Println()
}
