package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TodoList struct {
	ListName string  `json:ListName`
	List     *[]Todo `json:List`
}

func (tl *TodoList) addTodo(name string) Todo {
	created := time.Now().Local()
	return Todo{Name: name, Created: created}
}

func (tl *TodoList) listTodo() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Id", "Name", "Created"})
	for index, todo := range *tl.List {
		appendString := fmt.Sprintf("{%s, %s, %T}", index, todo.Name, todo.Created)
		table.Append(appendString)
	}
	table.Render()
}

func (tl *TodoList) remTodo(index int) {
	tl.List = slices.Delete(tl.List, index, index)
}

type Todo struct {
	Name    string    `json:Name`
	Created time.Time `json:Created`
}

const ConfigPath = ".config/todo"

func makeTodoFilePath(file string, tmp bool) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if tmp {
		return fmt.Sprintf("%s/%s/%s.json.tmp", home, ConfigPath, file), nil
	}

	return fmt.Sprintf("%s/%s/%s.json", home, ConfigPath, file), nil
}

func readTodoFile(file string) ([]byte, error) {
	todoFilePath, err := makeTodoFilePath(file, false)
	if err != nil {
		return nil, err
	}

	todoFile, err := os.ReadFile(todoFilePath)
	if err != nil {
		return nil, err
	}

	var data []byte
	err = json.Unmarshal(todoFile, &data)

	return data, err
}

func writeTodoFile(file string, data []byte) error {
	tmpTodoFilePath, err := makeTodoFilePath(file, true)
	if err != nil {
		return err
	}
	todoFilePath, err := makeTodoFilePath(file, false)
	if err != nil {
		return err
	}

	err = os.WriteFile(tmpTodoFilePath, data, 0666)
	if err != nil {
		return err
	}

	err = os.Rename(tmpTodoFilePath, todoFilePath)

	return err
}

func newTodoList(name string) TodoList {
	return TodoList{IdInc: 0, ListName: name, List: nil}
}
