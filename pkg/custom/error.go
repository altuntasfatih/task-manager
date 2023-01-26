package custom

import "errors"

var ErrRepairNotFound = errors.New("user not found")

type ErrorResponse struct {
	Message string `json:"message"`
}

func (c *ErrorResponse) Error() string {
	return c.Message
}
