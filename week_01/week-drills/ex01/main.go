// Écris une fonction `doubleAll(numbers *[]int)` qui multiplie 
// par 2 chaque élément d'un slice d'entiers, en modifiant l'original.

// Dans main :
// - Crée un slice []int{1, 2, 3, 4, 5}
// - Affiche-le avant
// - Appelle doubleAll
// - Affiche-le après

// Attendu :
// [1 2 3 4 5]
// [2 4 6 8 10]


package main

import "fmt"

func doubleAll(numbers *[]int) {
	for i, number := range *numbers {
		(*numbers)[i] = number * 2
	}
}

func display(numbers []int) { // No need to modify. So I don't need the poiter. The value is enough.

	fmt.Print("[")
	for i, number := range numbers {
		if i == len(numbers)-1 {
			fmt.Printf("%d", number)
		} else {
			fmt.Printf("%d ", number)
		}
	}
	fmt.Print("]\n")
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	display(numbers) // before
	doubleAll(&numbers)
	display(numbers) // after
}
