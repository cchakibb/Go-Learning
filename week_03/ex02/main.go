// Écris un programme qui simule le chargement de
// 5 fichiers en parallèle.

// Chaque goroutine :
// 1. Reçoit un nom de fichier (ex: "file1.txt", "file2.txt"...)
// 2. Simule un temps de chargement avec time.Sleep
//    (durée aléatoire entre 100ms et 500ms)
// 3. Affiche "Loaded: <filename> in <duration>"

// Main attend que tous les fichiers soient chargés,
// puis affiche "All files loaded."

// Tu auras besoin de :
// - time.Sleep(duration)
// - time.Duration
// - math/rand pour une durée aléatoire

package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

func chargeFile(i int) {
	defer wg.Done()
	duration := time.Duration(rand.Intn(401)+100) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Loaded file%d.txt in %d ms\n", i, duration.Milliseconds())

}

func main() {

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go chargeFile(i)
	}
	wg.Wait()
	fmt.Println("All files loaded")
}