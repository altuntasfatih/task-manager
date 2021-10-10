package models

import "time"

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type SetReminderRequest struct {
	Method ReminderMethod `json:"method" validate:"required,oneof=email onsite"`
}

type PeriodType string

const (
	Minute PeriodType = "minute"
	Hour   PeriodType = "hour"
	Day    PeriodType = "day"
)

type CreateTaskRequest struct {
	Name           string     `json:"name" validate:"required"`
	StartTime      time.Time  `json:"startTime" validate:"required"`
	EndTime        time.Time  `json:"endTime" validate:"required"`
	ReminderPeriod int64      `json:"reminderPeriod" validate:"required,gt=0"`
	PeriodType     PeriodType `json:"periodType" validate:"required,oneof=minute hour day"`
}
