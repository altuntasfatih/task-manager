package custom

type ErrorResponse struct {
	Message string `json:"message"`
}

func (c *ErrorResponse) Error() string {
	return c.Message
}
