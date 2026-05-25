// Dans main uniquement, pas de fonctions :

// 1. Déclare un slice []string vide
// 2. Ajoute 4 éléments avec append
// 3. Affiche la longueur avec len
// 4. Affiche uniquement les 2 derniers éléments avec un découpage
// 5. Modifie le premier élément directement via son index
// 6. Affiche le slice entier

package main

import "fmt"

func main() {

	s := []string{}

	s = append(s, "Go", "C", "C++", "Rust")
	fmt.Println(len(s))
	fmt.Println(s[2:])
	s[0] = "GoLang"
	fmt.Println(s)
}