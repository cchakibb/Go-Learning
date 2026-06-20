// 1. Crée store := &RoomStore{}
// 2. Enregistre les quatre routes avec le mux :
//    GET /rooms        → store.handleGetRooms
//    POST /rooms       → store.handleAddRoom
//    GET /rooms/{id}   → store.handleGetRoomByID
//    DELETE /rooms/{id}→ store.handleDeleteRoom
// 3. http.ListenAndServe(":8080", mux)

package main

import (
	"net/http"
)

func main() {
	store := &RoomStore{}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /rooms", store.handleGetRooms)
	mux.HandleFunc("POST /rooms", store.handleAddRoom)
	mux.HandleFunc("GET /rooms/{id}", store.handleGetRoomByID)
	mux.HandleFunc("DELETE /rooms/{id}", store.handleDeleteRoom)

	http.ListenAndServe(":8080", mux)
}


