// Dans main uniquement :

// 1. Déclare un slice []int avec 5 nombres
// 2. Crée un sous-slice des index 1 à 3 (inclus)
// 3. Modifie le premier élément du sous-slice
// 4. Affiche le slice original — que constates-tu ?
// 5. Affiche une explication en commentaire dans ton code

package main

import "fmt"

func main() {

	s := []int{5, 10, 15, 20, 25}
	s1 := s[1:4] // 10 - 15 - 20
	s1[0] = 999 // 10 -> 999 ===>> we modified the original slice, because we accessed it by index `[i]` 
	fmt.Println(s)
}