package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("os.UserConfigDir: %w", err)
	}

	configPath := fmt.Sprintf("%s/todo", dir)
	return configPath, nil

}

func MakeTodoFilePath(filename string, tmp bool) (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", fmt.Errorf("getConfigPath: %w", err)
	}

	if tmp {
		return fmt.Sprintf("%s/lists/%s.json.tmp", configPath, filename), nil
	}

	return fmt.Sprintf("%s/lists/%s.json", configPath, filename), nil
}

func ReadTodoFile(filename string) (List, error) {
	todoFilePath, err := MakeTodoFilePath(filename, false)
	if err != nil {
		return List{}, fmt.Errorf("MakeTodoFilePath: %w", err)
	}

	todoFile, err := os.ReadFile(todoFilePath)
	if err != nil {
		return List{}, fmt.Errorf("os.ReadFile: %w", err)
	}

	var todoList List
	err = json.Unmarshal(todoFile, &todoList)
	if err != nil {
		return List{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return todoList, nil
}

func WriteTodoFile(filename string, list List) error {
	tmpTodoFilePath, err := MakeTodoFilePath(filename, true)
	if err != nil {
		return fmt.Errorf("tmp MakeTodoFilePath: %w", err)
	}

	todoFilePath, err := MakeTodoFilePath(filename, false)
	if err != nil {
		return fmt.Errorf("MakeTodoFilePath: %w", err)
	}

	data, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = os.WriteFile(tmpTodoFilePath, data, 0666)
	if err != nil {
		return fmt.Errorf("")
	}

	err = os.Rename(tmpTodoFilePath, todoFilePath)

	return err
}
