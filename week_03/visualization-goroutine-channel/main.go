package main

import (
	"fmt"
)

var ch = make(chan int)

func calcSquare(i int) {
	square := i * i
	fmt.Printf("[Goroutine %d] J'ai calculé %d. J'essaie d'envoyer...\n", i, square)
	ch <- square
	fmt.Printf("[Goroutine %d] Libérée ! Mon score a été lu.\n", i)
}

func main() {
	for i := 1; i <= 2; i++ {
		go calcSquare(i)
	}

	for i := 1; i <= 2; i++ {
		fmt.Printf("[Main] J'attends une donnée à l'itération %d du bloc de lecture...\n", i)
		val := <-ch
		fmt.Printf("[Main] Reçu du canal : %d\n", val)
	}
}
