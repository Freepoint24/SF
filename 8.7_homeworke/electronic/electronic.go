package electronic

import "fmt"

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string // название ос
}

type applePhone struct {
	model string
	}

type androidPhone struct {
	brand string
	model string

}
type radioPhone struct {
	brand        string
	model        string

	ButtonsCount int
}

//четыре метода для Iphone
func (*applePhone) Brand() string	{ return "applePhone"}
func (p *applePhone) Model() string { return "p.model"
func (*applePhone) Type() string 	{return "smartphone"}
func (*applePhone) OS() string 		{return "IOS"}


// функция коструктор . NewApplePhone - maike iPhone
func NewApplePhone(model string() *applePhone{
	iphone := new(applePhone)
	iphone.model = model

	return iphon

}
