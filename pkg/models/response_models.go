package models

type GetUsersResponse struct {
	Users []*User
}
type GetUserResponse struct {
	*User
}

type GetTasksResponse struct {
	Tasks []*Task
}
type GetTaskResponse struct {
	*Task
}
