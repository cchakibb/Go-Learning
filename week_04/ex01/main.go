// Écris un programme avec deux channels :
// - ch1 envoie "source A" après 200ms (goroutine)
// - ch2 envoie "source B" après 100ms (goroutine)

// Utilise select pour afficher uniquement 
// la première réponse reçue, puis quitter.

// Attendu : "Première réponse : source B"

package main

import (
	"time"
	"fmt"
)


func sourceA(ch chan string) {

	time.Sleep(200 * time.Millisecond)

	ch <- "source A"
}

func sourceB(ch chan string) {

	time.Sleep(100 * time.Millisecond)

	ch <- "source B"
}


func main() {

	chA := make(chan string)
	chB := make(chan string)

	go sourceA(chA)
	go sourceB(chB)

	select {
	case valChA := <- chA:
		fmt.Println("First response: ", valChA)
	case valChaB := <- chB:
		fmt.Println("First response: ", valChaB)
	}
}