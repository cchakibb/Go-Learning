package main

import (
	"fmt"
)

func main() {

	tasks := []string{"Buy coffee", "Learn Go", "Fix bugs"}
	tasks = append(tasks, "commit", "push")
	fmt.Println(len(tasks)) // 5

	fmt.Println("----------")

	for i, task := range tasks {
		fmt.Println("Task #", i, task)
	}

	fmt.Println("----------")

	subTasks := tasks[0:3]
	for _, subTask := range subTasks {
		fmt.Println(subTask)
	}
	subTasks[0] = "Buy MOOOORE coffee"
	fmt.Println(tasks[0]) // original (tasks[0]) is modified. subTasks and tasks share memory /!\
}
