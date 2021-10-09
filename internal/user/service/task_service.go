package service

import (
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/store"
)

type TaskService interface {
	CreateTask(userId string, request *models.CreateTaskRequest) (*models.Task, error)
	GetTask(userId string, taskId int) (*models.Task, error)
	DeleteTask(userId string, taskId int) error
	GetTasks(userId string) ([]*models.Task, error)
}

func NewTaskService(userStore store.ReaderWriterRemover) (TaskService, error) {
	return &userService{
		userStore: userStore,
	}, nil
}

func (u *userService) CreateTask(userId string, request *models.CreateTaskRequest) (*models.Task, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}

	newTask := &models.Task{
		Id:                   len(user.Tasks) + 1,
		Name:                 request.Name,
		StartTime:            request.StartTime,
		EndTime:              request.EndTime,
		ReminderPeriodInHour: request.ReminderPeriodInHour,
	}

	if user.IsTaskOverLapWithOther(newTask) {
		return nil, custom.ErrTaskIsOverLap
	}

	user.AddTask(newTask)
	err = u.userStore.UpdateUser(user.Id, user)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}

func (u *userService) GetTasks(userId string) ([]*models.Task, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user.Tasks, nil
}

func (u *userService) GetTask(userId string, taskId int) (*models.Task, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}
	_, task, err := user.SearchTask(taskId)
	return task, err
}

func (u *userService) DeleteTask(userId string, taskId int) error {
	user, err := u.GetUser(userId)
	if err != nil {
		return err
	}
	index, _, err := user.SearchTask(taskId)
	if err != nil {
		return err
	}
	user.RemoveTask(index)
	return u.userStore.UpdateUser(user.Id, user)
}
