// Écris un programme qui :
// 1. Crée une map[string]int qui associe des prénoms à des âges
//    (au moins 4 entrées)
// 2. Écris une fonction `birthday(people map[string]int, name string) error`
//    qui incrémente l'âge d'une personne de 1.
//    Si la personne n'existe pas, retourne une erreur.
// 3. Dans main, appelle birthday sur un prénom qui existe
//    et un qui n'existe pas. Gère les deux cas.
// 4. Affiche la map avant et après.

package main

import (
	"errors"
	"fmt"
)

func birthday(people map[string]int, name string) error {

	_,ok := people[name]
	if ok {
		people[name]++
		return nil
	} 
	return errors.New("People not found")
}

func displayMap(people map[string]int) {

	for name,age := range people {
		fmt.Println(name, age)
	}
}

func main() {

	peopleAge := map[string]int{
		"Chakib": 21,
		"Naila":  20,
		"Lena":   5,
		"Luja":   3,
	}

	displayMap(peopleAge) // before
	fmt.Println("----")
	err := birthday(peopleAge, "Chakib") // Exists
	if err != nil {
		fmt.Print(err, "\n")
	}
	err = birthday(peopleAge, "John") // DOES NOT Exists
	if err != nil {
		fmt.Print(err, "\n")
	}
	displayMap(peopleAge) // After
}
