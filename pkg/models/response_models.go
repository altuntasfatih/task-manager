package models

type GetUsersResponse struct {
	Users []*User
}
type GetUserResponse struct {
	*User
}
