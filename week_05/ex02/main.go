// Ajoute une route POST /echo à ton serveur.
// Elle doit :
// 1. Lire le body de la requête avec io.ReadAll
// 2. Afficher dans le terminal (côté serveur) :
//    "Reçu X bytes: <contenu>"
// 3. Renvoyer le même contenu au client

// Test :
// curl -X POST http://localhost:8080/echo \
//      -d "hello server"

// Attendu côté client : hello server
// Attendu côté serveur : Reçu 12 bytes: hello server


package main

import (
	"net/http"
	"fmt"
	"io"
)

func handleHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from Go!\n")
}

func handleEcho(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
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
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found\n")
	})
	mux.HandleFunc("/echo", handleEcho)
	http.ListenAndServe(":8080", mux)
}