// Écris un programme avec :

// 1. Une fonction `keepDone(tasks []Task) []Task` qui prend un slice 
//    de Task et retourne un nouveau slice contenant uniquement 
//    les tâches où Done == true.
//    (Pas de modification en place ici — tu construis un nouveau slice)

// 2. Une struct Task avec Title string et Done bool.

// 3. Dans main :
//    - Crée un slice de 5 tasks, dont 2 marquées Done: true
//    - Appelle keepDone
//    - Affiche le slice original (doit être intact)
//    - Affiche le slice retourné (doit contenir uniquement les done)


package main

import "fmt"

type Task struct {
	Title	string
	Done	bool
}

func keepDone(tasks []Task) []Task {

	newSlice := []Task{}
	for _, task := range tasks {
		if task.Done {
			newSlice = append(newSlice, task)
		}
	}
	return newSlice
}

func displaySlice (tasks []Task) {
	
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func main() {

	s1 := []Task{{Title: "Eat", Done: false}, {Title: "Sleep", Done: false}, {Title: "Learn Go", Done: true}, 
		{Title: "Practice", Done: true}, {Title: "Forget C", Done: false}}
	s2 := keepDone(s1)
	displaySlice(s1)
	fmt.Println()
	displaySlice(s2)
}	

