package main

import "interF/models"

func main() {
	ownAlarm := models.OwnAlarm{
		ID:             1,
		WelcomeMassage: "Lalala",
	}
	alarmExample(ownAlarm)
}

func alarmExample(alarm models.OwnAlarm) {
	alarm.PlayAlarm()
}
