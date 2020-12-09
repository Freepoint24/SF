package smplStr

//вычисление площади прямоугольника
// структура kvadrat c полями storona1, storona1, square тип float64
type rectangle struct {
	storona1 float64
	storona2 float64
	square   float64
}

// инициализиру структуру rectangle функцией NewKvadrat
// return возвращает копию структуры rectangle
func NewKvadrat() rectangle {
	return rectangle{}
}

// метод Kvadr получает значения storona1 и storona2 структуры rectangle из функции NewKvadrat
// b return возращает значение square умножение storona1 на storona2
func (x *rectangle) Kvadr(storona1, storona2 float64) (square float64) {
	return storona1 * storona2
}
