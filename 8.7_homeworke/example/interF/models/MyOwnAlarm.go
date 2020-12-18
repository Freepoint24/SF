package models

import "time"

type OwnAlarm struct {
	ID int
	WelcomeMassage string
	playDate time.Time
}

func (a *OwnAlarm) setAlarm(s string) {
	panic("implement me")
}

func (a *OwnAlarm) setTime(time.Time) {
	panic("implement me")
}

func (a *OwnAlarm) PlayAlarm() {
	panic("implement me")
}

