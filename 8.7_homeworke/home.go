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
	Brand string
	Model string
	Type  string
	OS    string
}
type androidPhone struct {
	Brand string
	Model string
	Type  string
	OS    string
}
type radioPhone struct {
	Brand        string
	Model        string
	Type         string
	ButtonsCount int
}

func (a *Smartphone) Type() string {
	return fmt.Sprintf( " "	, a.applePhone, a.Smartphone)
}
