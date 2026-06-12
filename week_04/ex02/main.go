// Même setup, mais cette fois utilise select 
// avec un timeout : si aucune source ne répond 
// en moins de 150ms, affiche "timeout" et quitte.

// Indice : time.After(duration) retourne un channel 
// qui envoie une valeur après la durée spécifiée.

// Teste avec :
// - sourceA : 200ms, sourceB : 180ms (→ timeout)
// - sourceA : 200ms, sourceB : 100ms (→ source B gagne)

package main

import (
	"time"
	"fmt"
)


func sourceA(ch chan string) {

	time.Sleep(140 * time.Millisecond)

	ch <- "source A"
}

func sourceB(ch chan string) {

	time.Sleep(210 * time.Millisecond)

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
	case <- time.After(150 * time.Millisecond): // if no source responds before `150ms`
		fmt.Println("timeout")
	}
}