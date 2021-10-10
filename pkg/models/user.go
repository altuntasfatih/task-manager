package models

import (
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"sort"
)

type User struct {
	Id             string         `json:"id"`
	Email          string         `json:"email"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Tasks          TaskList       `json:"tasks"`
	ReminderMethod ReminderMethod `json:"reminderMethod"`
}
type ReminderMethod string

const (
	Email  ReminderMethod = "email"
	Onsite ReminderMethod = "onsite"
)

func NewUser(id, email, firsName, lastName string) *User {
	return &User{Id: id, Email: email, FirstName: firsName, LastName: lastName, Tasks: make([]*Task, 0)}
}

func (u *User) SortTasks() {
	sort.Sort(u.Tasks)
}

func (u *User) AddTask(task *Task) error {
	if u.IsTaskOverLapWithOther(task) {
		return custom.ErrTaskIsOverLap
	}
	u.Tasks = append(u.Tasks, task)
	u.SortTasks()
	return nil
}

func (u *User) SearchTask(taskId int) (index int, task *Task, err error) {
	for i, task := range u.Tasks {
		if task.Id == taskId {
			return i, task, nil
		}
	}
	return 0, nil, custom.ErrTaskNotFound
}

func (u *User) RemoveTask(taskId int) error {
	index, _, err := u.SearchTask(taskId)
	if err != nil {
		return err
	}
	u.Tasks = append(u.Tasks[:index], u.Tasks[index+1:]...)
	return nil
}

func (u User) IsTaskOverLapWithOther(newTask *Task) bool {
	taskLength := u.Tasks.Len()
	if taskLength == 0 {
		return false
	}

	lastIndex := sort.Search(taskLength, func(i int) bool {
		return u.Tasks[i].StartTime.After(newTask.StartTime)
	})

	if lastIndex == 0 {
		return u.Tasks[0].EndTime.After(newTask.EndTime)
	}

	if lastIndex == taskLength {
		return u.Tasks[lastIndex-1].EndTime.After(newTask.StartTime)

	}
	return u.Tasks[lastIndex].EndTime.After(newTask.StartTime)
}

type TaskList []*Task

func (tasks TaskList) Len() int {
	return len(tasks)
}

func (tasks TaskList) Less(i, j int) bool {
	return tasks[i].StartTime.Before(tasks[j].StartTime)
}

func (tasks TaskList) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}
