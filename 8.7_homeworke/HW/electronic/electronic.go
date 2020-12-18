package electronic

//интерфесы
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

//структуры
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
	buttonsCount int
}

//методы
//методы для applePhone
func (*applePhone) Brand() string   { return "applePhone" }
func (p *applePhone) Model() string { return p.model }
func (*applePhone) Type() string    { return "smartphone" }
func (*applePhone) OS() string      { return "iOS" }

//методы для  androidPhone
func (p *androidPhone) Brand() string { return p.brand }
func (p *androidPhone) Model() string { return p.model }
func (p *androidPhone) Type() string  { return "smartphone" }
func (p *androidPhone) OS() string    { return "Android" }

//методы для radioPhone
func (p *radioPhone) Brand() string     { return p.brand }
func (p *radioPhone) Model() string     { return p.model }
func (p *radioPhone) Type() string      { return "station" }
func (p *radioPhone) ButtonsCount() int { return p.buttonsCount }

//функции-конструкторы
//для структуры applePhone
func NapplePhone(model string) *applePhone {
	iphone := new(applePhone)
	iphone.model = model
	return iphone
}

//для структуры androidPhone
func NandroidPhone(brand, model string) *androidPhone {
	android := new(androidPhone)
	android.brand = brand
	android.model = model
	return android
}

//для структуры radioPhone
func NradioPhone(brand string, model string, buttonsCount int) *radioPhone {
	radio := new(radioPhone)
	radio.brand = brand
	radio.model = model
	radio.buttonsCount = buttonsCount
	return radio
}
