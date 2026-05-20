package main

import "fmt"

func main() {

	taskStatus := map[string]bool{
		"study": true,
		"sleep": false,
		"drink": false,
		"eat":   true,
	}

	for task, done := range taskStatus {
		fmt.Println(task, done)
	}

	fmt.Println("--------------")

	val, okay := taskStatus["study"]
	if okay {
		fmt.Println("found:", val)
	} else {
		fmt.Println("Not found")
	}

	val, okay = taskStatus["Non existing task"]
	if okay {
		fmt.Println("found:", val)
	} else {
		fmt.Println("Not found")
	}

	delete(taskStatus, "eat")

	fmt.Println("--------------")

	for task, done := range taskStatus {
		fmt.Println(task, done) // 'eat' was deleted
	}
}
