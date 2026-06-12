// Écris un programme qui :
// 1. Lit une variable d'environnement "ROOM_ID"
//    Si elle n'existe pas, utilise "101" comme valeur par défaut
// 2. Écrit dans un fichier "room.txt" la ligne :
//    "Checking room: <ROOM_ID>"
// 3. Relit ce fichier et affiche son contenu

package main

import (
	"os"
	"fmt"
)


func main() {

	roomIdEnv := os.Getenv("ROOM_ID")
	if roomIdEnv == "" {
		roomIdEnv = "101"
	}

	content := fmt.Sprintf("Checking room: %s", roomIdEnv)
	errWriteFile := os.WriteFile("room.txt", []byte(content), 0644)
	if errWriteFile != nil {
		fmt.Println(errWriteFile)
	}
	

	data, errReadFile := os.ReadFile("room.txt")
	if errReadFile != nil {
		fmt.Println(errReadFile)
	} else {
		fmt.Printf("%s\n", data)
	}
}