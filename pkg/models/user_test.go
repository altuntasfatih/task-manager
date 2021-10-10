package models

import (
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUser_SortTasks(t *testing.T) {

	task1 := NewTask(1, "task1", time.Now(), time.Now().Add(1*time.Hour), 5, Minute)
	task2 := NewTask(2, "task2", time.Now().Add(-9*time.Hour), time.Now().Add(-8*time.Hour), 5, Minute)
	task3 := NewTask(3, "task3", time.Now().Add(4*time.Hour), time.Now().Add(5*time.Hour), 5, Minute)
	task4 := NewTask(4, "task4", time.Now().Add(-3*time.Hour), time.Now().Add(-2*time.Hour), 5, Minute)
	task5 := NewTask(5, "task5", time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour), 5, Minute)
	user := &User{
		Tasks: TaskList{task1, task2, task3, task4, task5},
	}
	user.SortTasks()

	require.Equal(t, user.Tasks[0], task2)
	require.Equal(t, user.Tasks[1], task4)
	require.Equal(t, user.Tasks[2], task1)
	require.Equal(t, user.Tasks[3], task5)
	require.Equal(t, user.Tasks[4], task3)
}
func TestUser_IsTaskOverLapWithOther(t *testing.T) {

	now := time.Now()
	task1 := NewTask(1, "task1", now, now.Add(10*time.Minute), 5, Minute)
	task2 := NewTask(2, "task2", now.Add(15*time.Minute), now.Add(25*time.Minute), 5, Minute)
	task3 := NewTask(3, "task3", now.Add(30*time.Minute), now.Add(40*time.Minute), 5, Minute)
	task4 := NewTask(4, "task4", now.Add(45*time.Minute), now.Add(55*time.Minute), 5, Minute)
	task5 := NewTask(5, "task5", now.Add(60*time.Minute), now.Add(70*time.Minute), 5, Minute)

	taskList := TaskList{task1, task2, task3, task4, task5}
	user := &User{}

	for _, task := range taskList {
		err := user.AddTask(task)
		require.Nil(t, err)
	}

	task6 := NewTask(6, "task6", now.Add(-10*time.Minute), now.Add(1*time.Minute), 6, Minute)
	task7 := NewTask(7, "task7", now.Add(28*time.Minute), now.Add(33*time.Minute), 7, Minute)
	task8 := NewTask(8, "task8", now.Add(69*time.Minute), now.Add(85*time.Minute), 8, Minute)
	overlappingTaskList := TaskList{task6,
		task7,
		task8,
		task3,
		task1,
	}

	for _, task := range overlappingTaskList {
		err := user.AddTask(task)
		require.Equal(t, err, custom.ErrTaskIsOverLap)
	}
}
