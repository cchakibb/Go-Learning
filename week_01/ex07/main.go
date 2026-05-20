package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Title string
	Done  bool
}

func findTask(tasks []Task, title string) (Task, error) {

	for _, task := range tasks {
		if task.Title == title {
			return task, nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func main() {

	tasks := []Task{{Title: "Step 1", Done: false}, {Title: "Step 2", Done: false}, {Title: "Step 3", Done: false}}

	foundTask, err := findTask(tasks, "Step 1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task found:", foundTask.Title)
	}

	foundTask, err = findTask(tasks, "Non existing task")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task found:", foundTask.Title)
	}
}
