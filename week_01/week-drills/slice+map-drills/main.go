// Dans main uniquement :

// 1. Déclare une map[string][]int
// 2. Ajoute directement 3 entrées avec des slices de nombres
// 3. Ajoute un nombre à l'une des entrées existantes avec append
// 4. Affiche toute la map

package main

import "fmt"

func main() {

		m := make(map[string][]int)
		m["temperature"] = []int{24, 31, 16}
		m["grades"] = []int{18, 12, 9}
		m["age"] = []int{5, 30, 42}
	
		// m := map[string][]int {
		// 	"temperature": {24, 31, 16},
		// 	"grades": {18, 12, 9},
		// 	"age": {5, 30, 42},
		// }

		m["temperature"] = append(m["temperature"], 500)
		fmt.Println(m)
}