package custom

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrTaskNotFound = errors.New("task not found")
var ErrTaskIsOverLap = errors.New("task is overlapping with another task")

type ErrorResponse struct {
	Message string `json:"message"`
}

func (c *ErrorResponse) Error() string {
	return c.Message
}
