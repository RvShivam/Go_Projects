package internal

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)

	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil // file not created yet
	}

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func AddTask(title string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	id := len(tasks) + 1

	task := Task{
		ID:        id,
		Title:     title,
		Status:    "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, task)
	return saveTasks(tasks)
}

func UpdateTask(id int, newTitle string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = newTitle
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return saveTasks(tasks)
}

func UpdateTaskStatus(id int, newStatus string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return errors.New("task not found")
	}

	return saveTasks(tasks)
}

func ListTasks() ([]Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	newTasks := []Task{}

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		return errors.New("task not found")
	}

	return saveTasks(newTasks)
}
