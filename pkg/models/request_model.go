package models

import "time"

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type CreateTaskRequest struct {
	Name                 string    `json:"name" validate:"required"`
	StartTime            time.Time `json:"startTime" validate:"required"`
	EndTime              time.Time `json:"endTime" validate:"required"`
	ReminderPeriodInHour int       `json:"reminderPeriodInHour" validate:"required,gt=0"`
}
