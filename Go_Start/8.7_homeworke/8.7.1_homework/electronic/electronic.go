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
func (x *applePhone) Model() string { return x.model }
func (*applePhone) Type() string    { return "smartphone" }
func (*applePhone) OS() string      { return "iOS" }

//методы для  androidPhone
func (x *androidPhone) Brand() string { return x.brand }
func (x *androidPhone) Model() string { return x.model }
func (x *androidPhone) Type() string  { return "smartphone" }
func (x *androidPhone) OS() string    { return "Android" }

//методы для radioPhone
func (x *radioPhone) Brand() string     { return x.brand }
func (x *radioPhone) Model() string     { return x.model }
func (x *radioPhone) Type() string      { return "station" }
func (x *radioPhone) ButtonsCount() int { return x.buttonsCount }

//функции-конструкторы:
//для структуры applePhone
func NapplePhone(model string) *applePhone {
	iphone := new(applePhone) //указатель на структуру applePhone
	iphone.model = model
	return iphone
}

//для структуры androidPhone
func NandroidPhone(brand, model string) *androidPhone {
	android := new(androidPhone) //указатель на структуру androidPhone
	android.brand = brand
	android.model = model
	return android
}

//для структуры radioPhone
func NradioPhone(brand string, model string, buttonsCount int) *radioPhone {
	radio := new(radioPhone) //указатель на структуру radioPhone
	radio.brand = brand
	radio.model = model
	radio.buttonsCount = buttonsCount
	return radio
}
