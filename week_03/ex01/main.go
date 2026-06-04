// Écris un programme qui lance 5 goroutines.
// Chaque goroutine reçoit un numéro (1 à 5) et affiche :
// "Goroutine X démarre" puis "Goroutine X termine"

// Utilise un WaitGroup pour que main attende
// que toutes les goroutines aient terminé.

package main

import ("sync"
		"fmt")


var wg sync.WaitGroup

func goroutine(i int) {
	defer wg.Done()
	fmt.Printf("Goroutine %d starts\n", i)
	fmt.Printf("Goroutine %d ends\n", i)
}


func main() {


	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go goroutine(i)
	}
	wg.Wait()
	
}
