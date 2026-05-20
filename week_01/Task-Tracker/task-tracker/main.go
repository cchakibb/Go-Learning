package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

// Commands

func addTask(task *[]Task, title string) {

	if len(strings.TrimSpace(title)) == 0 { // Do not accept `add(space) ` 
		fmt.Println("Please provide a valid task name")
		return
	}
	t := Task{Title: title, Done: false}

	*task = append(*task, t)
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add a task first: add <task>")
		return
	}

	for _, task := range tasks {
		if task.Done == true {
			fmt.Printf("[✓] %s\n", task.Title)
		} else {
			fmt.Printf("[ ] %s\n", task.Title)
		}
	}
}

func doneTask(tasks *[]Task, title string) error {

	for i, taskDone := range *tasks {
		if taskDone.Title == title {
			(*tasks)[i].Done = true
			return nil
		}
	}
	return errors.New("Task Tracker: Task not found")
}

func deleteTask(tasks *[]Task, title string) error {
	for i, taskToDelete := range *tasks {
		if taskToDelete.Title == title {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return nil
		}
	}
	return errors.New("Task Tracker: Task not found")
}

func main() {

	tasks := []Task{}

	fmt.Println("Task Tracker. Commands: add <title>, list, done <title>, delete <title>, quit")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		usrInput := strings.SplitN(line, " ", 2) // add <task> ...

		command := usrInput[0] // Task commands: add, list, delete, quit

		switch command {
		case "add":
			if len(usrInput) < 2 {
				fmt.Println("usage: add <title>")
			} else {
				addTask(&tasks, usrInput[1])
			}
		case "list":
			listTasks(tasks)
		case "delete":
			if len(usrInput) < 2 {
				fmt.Println("usage: delete <title>")
			} else {
				err := deleteTask(&tasks, usrInput[1])
				if err != nil {
					fmt.Println(err)
				}
			}
		case "done":
			if len(usrInput) < 2 {
				fmt.Println("usage: done <title>")
			} else {
				err := doneTask(&tasks, usrInput[1])
				if err != nil {
					fmt.Println(err)
				}
			}
		case "quit":
			return
		default:
			fmt.Println("This command does not exist. You can use: add <task>, list, delete <task> , quit")
		}
	}
}
