package main

import (
	"module1/test"
)

func  main()  {
	// Никаких проблем с вызовом функции
	test.Public()

	// При запуск программы выдаст ошибку
	test.Private()


}
