// Écris un programme qui :
// 1. Lance 5 goroutines
// 2. Chaque goroutine calcule le carré de son numéro (i * i)
//    et l'envoie dans un channel
// 3. Main reçoit les 5 résultats depuis le channel
//    et les affiche
// 4. Pas de WaitGroup cette fois — le channel suffit pour synchroniser

package main

import ("fmt"
)
		

func calcSquare(i int, ch chan int){

	square := i * i
	ch <- square
	
}

func main() {

	ch := make(chan int)

	for i := 1; i <= 5 ; i++ {
		go calcSquare(i, ch)
	}

	for i := 1; i <= 5; i++ {
		val := <- ch
		fmt.Println(val)
	}
}