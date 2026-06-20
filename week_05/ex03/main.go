// Ajoute deux routes à ton serveur :

// GET /room
// → Retourne une Room hardcodée en JSON
//   {"id":"101","name":"Suite","price":150}

// POST /room
// → Reçoit une Room en JSON dans le body
// → Affiche dans le terminal les champs décodés
// → Retourne {"status":"created"}

// Utilise json.NewEncoder et json.NewDecoder.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Room struct {
	ID 		string 		`json:"id"`
	Name 	string 		`json:"name"`
	Price 	float64 	`json:"price"`
}

func handleHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from Go!\n")
}

func handleGetRoom(w http.ResponseWriter, r *http.Request){
	room := Room{ID: "101", Name: "Suite", Price: 150.0}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(room)

}

func handlePostRoom(w http.ResponseWriter, r *http.Request){
	var room Room
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&room)
	fmt.Printf("id: %s\nname: %s\nprice: %f", room.ID, room.Name, room.Price)
	w.Write([]byte(`{"status":"created"}` + "\n"))
}

func handleEcho(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading body %s", err)
		return
	}
	fmt.Printf("Received %d bytes: %s\n", len(body), body)
	fmt.Fprintf(w, "%s\n", body)
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("GET /room", handleGetRoom)
	mux.HandleFunc("POST /room", handlePostRoom)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found\n")
	})
	mux.HandleFunc("/echo", handleEcho)
	http.ListenAndServe(":8080", mux)
}