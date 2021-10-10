package service

import (
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/storage"
)

type TaskService interface {
	CreateTask(userId string, request *models.CreateTaskRequest) (*models.Task, error)
	GetTask(userId string, taskId int) (*models.Task, error)
	DeleteTask(userId string, taskId int) error
	GetTasks(userId string) ([]*models.Task, error)
}

func NewTaskService(userStore storage.ReaderWriterRemover) (TaskService, error) {
	return &userService{
		userStore: userStore,
	}, nil
}

func (u *userService) CreateTask(userId string, request *models.CreateTaskRequest) (*models.Task, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}

	newTask := models.NewTask(len(user.Tasks)+1, request.Name, request.StartTime, request.EndTime, request.ReminderPeriod, request.PeriodType)

	if err := user.AddTask(newTask); err != nil {
		return nil, err
	}

	if err := u.userStore.UpdateUser(user.Id, user); err != nil {
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
	if err := user.RemoveTask(taskId); err != nil {
		return err
	}
	return u.userStore.UpdateUser(user.Id, user)
}
