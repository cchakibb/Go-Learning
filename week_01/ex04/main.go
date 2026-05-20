package main

import "fmt"

type Task struct {
	Title    string
	Done     bool
	Priority int
}

// Methods
func (t Task) Display() {
	fmt.Println("- Task name:", t.Title, "- Done:", t.Done, "- Priority:", t.Priority)
}
func (t *Task) Complete() {
	t.Done = true
}

func main() {
	task1 := Task{Title: "Buy coffee", Done: false, Priority: 5}
	task2 := Task{Title: "Take a nap", Done: false, Priority: 1}
	task1.Display()
	task2.Display()
	task1.Complete()
	task1.Display()
	task2.Display()
}
