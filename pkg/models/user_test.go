package models

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUser_SortTasks(t *testing.T) {

	task1 := NewTask(1, "task1", time.Now(), time.Now().Add(1*time.Hour), 5)
	task2 := NewTask(2, "task2", time.Now().Add(-9*time.Hour), time.Now().Add(-8*time.Hour), 5)
	task3 := NewTask(3, "task3", time.Now().Add(4*time.Hour), time.Now().Add(5*time.Hour), 5)
	task4 := NewTask(4, "task4", time.Now().Add(-3*time.Hour), time.Now().Add(-2*time.Hour), 5)
	task5 := NewTask(5, "task5", time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour), 5)
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
