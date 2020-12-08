package smplStr

// структура kvadrat c полями storona1, storona1, square тип float64
type kvadrat struct {
	storona1 float64
	storona2 float64
	square float64
}

func NewKvadrat() kvadrat  {
	return kvadrat{}
}

func (x *kvadrat) Kvadr(storona1, storona2 float64) (square float64)   {
	square = x.storona1 * x.storona2
	return
}

func (x *kvadrat) add(storona1, storona2 float64) float64 {
	return  storona1 * storona2

}