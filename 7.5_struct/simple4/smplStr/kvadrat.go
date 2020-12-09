package smplStr

// структура kvadrat c полями storona1, storona1, square тип float64
type kvadrat struct {
	storona1 float64
	storona2 float64
	square   float64
}
// инициализиру структуру kvadrat функцией NewKvadrat
// return возвращает копию структуры kvadrat
func NewKvadrat() kvadrat {
	return kvadrat{}
}

// метод Kvadr получает значения storona1 и storona2 структуры kvadrat
//из функции NewKvadrat и возвращает умножение storona1 на storona2
func (x *kvadrat) Kvadr(storona1, storona2 float64) (square float64) {
//square = x.storona1 * x.storona2
	return storona1 * storona2
}

//func (x *kvadrat) add(storona1, storona2 float64) float64 {
//	return  storona1 * storona2

//}
