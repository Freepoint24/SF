package main

import "fmt"

func main() {
	// По умолчанию будет тип int
	i := -5
	fmt.Println(i)

	// Явно указываем, что нам нужен тип int16
	var i16 int16 = 6
	fmt.Println(i16)

	var ui8 uint8 = 7
	var b byte = 7
	// Убедились, что byte - это алиас на тип uint8
	if ui8 == b {
		fmt.Println(fmt.Sprintf("%d и %d равны", b))
	}

	// Беззнаковое целое число. Данной переменной нельзя присвоить отрицательное занчение
	var ui uint = 1
	fmt.Println(ui)

	// Число с плавающей запятой (по умолчанию float64)
	f := 3.2
	fmt.Println(f)

	// Явно задаем тип float32
	var f32 float32 = 3.3
	fmt.Println(f32)

	// Логический тип. Может принимать только значение true | false
	yes := true
	fmt.Println(yes)

	var no = false
	fmt.Println(no)

}
