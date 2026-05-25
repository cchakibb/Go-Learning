// Écris un programme qui :

// 1. Crée une map[string][]string qui associe une catégorie 
//    à une liste de tâches.
//    Exemple :
//    "work"     → ["emails", "meeting", "report"]
//    "personal" → ["sport", "groceries"]
//    "learning" → ["Go", "algorithms"]

// 2. Écris une fonction `addToCategory` qui prend :
//    - la map
//    - une catégorie (string)
//    - une tâche (string)
//    et ajoute la tâche au slice de cette catégorie.
//    Si la catégorie n'existe pas, elle la crée.

// 3. Écris une fonction `displayAll` qui affiche toutes 
//    les catégories et leurs tâches.

// 4. Dans main :
//    - Initialise la map avec les données ci-dessus
//    - Affiche tout
//    - Ajoute "testing" à "work" et "piano" à "personal"
//    - Ajoute une nouvelle catégorie "health" avec la tâche "sleep"
//    - Affiche tout à nouveau


package main

import "fmt"

func addToCategory(m map[string][]string, category string, task string) {
	
	val, _ := m[category]
		m[category] = append(val, task)
}


func displayAll(m map[string][]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}


func main() {

	m := map[string][]string {
		"work": {"emails", "meeting", "report"}, 
		"personal": {"sport", "groceries"},
		"learning": {"Go", "algorithms"},
	}
	displayAll(m)
	fmt.Println()
	addToCategory(m, "work", "testing")
	addToCategory(m, "personal", "piano")
	addToCategory(m, "health", "sleep")
	displayAll(m)
}