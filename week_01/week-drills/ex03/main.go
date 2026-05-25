// Crée une struct `Counter` avec un seul champ `Value int`.

// Écris deux méthodes :
// - `Increment()` qui ajoute 1 au counter
// - `Display()` qui affiche la valeur

// Dans main :
// - Crée un Counter à 0
// - Appelle Increment() 3 fois
// - Appelle Display()
// - Attendu : 3

// Choisis consciemment le bon type de receiver pour chaque méthode.
// Si le résultat affiche 0, tu as fait le mauvais choix quelque part.

package main

import (
	"fmt"
)

// struct
type Counter struct {

	Value	int
}

// Methods
func (c *Counter) increment () { // Pointer since we are modifying
	c.Value++
}

func (c Counter) display () { // Value. No modification, no pointer

	fmt.Println(c.Value)
}

func main() {

	c := Counter{Value: 0}
	c.increment()
	c.increment()
	c.increment()
	c.display()
}