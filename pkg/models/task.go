package models

import "time"

type Task struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	ReminderPeriod int64     `json:"reminderPeriod"` //in nanosecond
}

func NewTask(id int, name string, start, end time.Time, reminderPeriod int64, periodType PeriodType) *Task {
	calculatePeriod := func(period int64, periodType PeriodType) int64 {
		switch periodType {
		case Minute:
			return period * time.Minute.Nanoseconds()
		case Hour:
			return period * time.Hour.Nanoseconds()
		case Day:
			return period * time.Hour.Nanoseconds() * 24
		default:
			return period * time.Minute.Nanoseconds()
		}
	}
	return &Task{Id: id, Name: name, StartTime: start, EndTime: end, ReminderPeriod: calculatePeriod(reminderPeriod, periodType)}
}
