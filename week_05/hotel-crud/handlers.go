// 1. handleGetRooms		→ GET /rooms			→ retourne toutes les rooms en JSON
// 2. handleAddRoom			→ POST /rooms			→ décode le body, ajoute, retourne la room créée
// 3. handleGetRoomByID		→ GET /rooms/{id}		→ retourne une room, 404 si introuvable
// 4. handleDeleteRoom		→ DELETE /rooms/{id}	→ supprime, 404 si introuvable

package main

import (
	"encoding/json"
	"net/http"
)

func (s *RoomStore) handleGetRooms(w http.ResponseWriter, _ *http.Request) {
	rooms := s.GetAll()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(rooms)
}

func (s *RoomStore) handleAddRoom(w http.ResponseWriter, r *http.Request) {
	var room Room

	decoder := json.NewDecoder(r.Body)
	if err:= decoder.Decode(&room); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.Add(room)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(room)
}

func (s *RoomStore) handleGetRoomByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	room, err := s.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func (s *RoomStore) handleDeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	room, err := s.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}