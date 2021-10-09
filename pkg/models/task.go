package models

import "time"

type Task struct {
	Id                   int
	Name                 string
	StartTime            time.Time
	EndTime              time.Time
	ReminderPeriodInHour int
}

func NewTask(id int, name string, start, end time.Time, reminder int) *Task {
	return &Task{Id: id, Name: name, StartTime: start, EndTime: end, ReminderPeriodInHour: reminder}
}
