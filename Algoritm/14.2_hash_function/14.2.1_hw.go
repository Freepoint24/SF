package main

import "fmt"

//1. отрицательных значений при преобразовании строки в байты не бывает
//2. юинт в два раза более емкий, чем инт

func hashint64(val int64) int64 {
	return (val % 1000)
}

func main() {
	y := 564564654
	x := hashint64(int64(y))
	fmt.Println(x)

}
