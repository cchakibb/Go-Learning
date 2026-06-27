package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Chaîne de connexion : toutes les infos pour joindre PostgreSQL dans Docker
	connStr := "host=localhost port=5433 user=hotel password=secret dbname=hoteldb sslmode=disable"

	// sql.Open prépare la connexion mais ne se connecte pas encore
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err) // si la config est invalide, on quitte immédiatement
	}
	defer db.Close() // ferme la connexion quand main() se termine

	// db.Ping() établit la vraie connexion — c'est ici qu'on sait si PostgreSQL répond
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// Crée le store — l'objet qui sait parler à la base
	store := NewPostgresStore(db)

	// Déclare les routes HTTP
	mux := http.NewServeMux()
	mux.HandleFunc("GET /rooms", store.handleGetAll)
	mux.HandleFunc("POST /rooms", store.handleAdd)
	mux.HandleFunc("GET /rooms/{id}", store.handleGetByID)
	mux.HandleFunc("DELETE /rooms/{id}", store.handleDelete)
	mux.HandleFunc("PATCH /rooms/{id}/price", store.handleUpdatePrice)

	// Démarre le serveur — bloque jusqu'à arrêt
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
