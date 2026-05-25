// 1. Déclare une map[string]int vide avec make
// 2. Ajoute 3 entrées
// 3. Affiche la valeur d'une clé qui existe avec value, ok
// 4. Affiche la valeur d'une clé qui n'existe pas avec value, ok

package main

import "fmt"

func main() {

	m := make(map[string]int)
		m["one"] = 1
		m["two"] = 2
		m["three"] = 3

/*
	m := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}
*/

	value, ok := m["one"]
	if ok {
		fmt.Println(value)
	}

	value, ok = m["nine"]
	if ok {
		fmt.Println(value)
	}
}