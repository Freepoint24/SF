package models

import "time"

type Alarm interface {
	setTime(time.Time)
	setAlarm(string)
	PlayAlarm()
}
