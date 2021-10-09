package service

import (
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
	"time"
)

func Test_Sort(t *testing.T) {
	task1 := models.NewTask(1, "task1", time.Now(), time.Now().Add(1*time.Hour), 5)
	task2 := models.NewTask(2, "task2", time.Now().Add(-9*time.Hour), time.Now().Add(-8*time.Hour), 5)
	task3 := models.NewTask(3, "task3", time.Now().Add(4*time.Hour), time.Now().Add(5*time.Hour), 5)
	task4 := models.NewTask(4, "task4", time.Now().Add(-3*time.Hour), time.Now().Add(-2*time.Hour), 5)
	task5 := models.NewTask(5, "task5", time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour), 5)
	taskList := models.TaskList{task1, task2, task3, task4, task5}
	sort.Sort(taskList)

	require.Equal(t, taskList[0], task2)
	require.Equal(t, taskList[1], task4)
	require.Equal(t, taskList[2], task1)
	require.Equal(t, taskList[3], task5)
	require.Equal(t, taskList[4], task3)
}
