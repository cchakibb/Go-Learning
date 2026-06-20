// Écris un serveur qui :
// 1. Écoute sur le port 8080
// 2. Répond "Hello from Go!" sur la route GET /hello
// 3. Répond "Not found" avec un status 404 sur toute autre route

// Pour lancer : go run main.go
// Pour tester : curl http://localhost:8080/hello
//               curl http://localhost:8080/autre

package main

import (
	"net/http"
	"fmt"
)

func maFonction(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from Go!\n")
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", maFonction)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found\n")
	})
	http.ListenAndServe(":8080", mux)
}