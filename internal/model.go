package internal

import (
	"fmt"
	"slices"
	"time"

	"github.com/spf13/viper"
)

type List struct {
	Name  string `json:"ListName"`
	Tasks []Task `json:"ListTasks"`
}

type Task struct {
	Name    string `json:"TaskName"`
	Created string `json:"TaskCreated"`
}

func (l *List) AddTask(task string) {
	current_time := time.Now().Local()
	created := fmt.Sprintf("%d-%02d-%02d %02d:%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute())
	todo := Task{Name: task, Created: created}
	l.Tasks = append(l.Tasks, todo)
}

func (l *List) RemoveTask(index int) {
	l.Tasks = slices.Delete(l.Tasks, index, index)
}

func (tk *Task) ModTodo(index int, task string) {

}

func NewList(name string) List {
	return List{Name: name, Tasks: nil}
}

func GetActiveList() (List, error) {
	activeTodo := viper.GetString("ActiveTodo")
	todoList, err := ReadTodoFile(activeTodo)
	if err != nil {
		return List{}, err
	}

	return todoList, nil
}
