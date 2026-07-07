package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s *PostgresStore) handleAdd(w http.ResponseWriter, r *http.Request) {
	var room Room

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&room); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := s.Add(room); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

func (s *PostgresStore) handleGetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3 * time.Second)
	defer cancel()
	rooms, err := s.GetAll(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout) // 504
			return
		}
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

func (s *PostgresStore) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	room, err := s.GetByID(id)
	if errors.Is(err, ErrRoomNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func (s *PostgresStore) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := s.Delete(id)
	if errors.Is(err, ErrRoomNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *PostgresStore) handleUpdatePrice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Struct anonyme pour décoder {"price": 299.99} depuis le body
	var body struct {
		Price float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Appelle UpdatePrice et gère les erreurs
	if err := s.UpdatePrice(id, body.Price); err != nil {
		if errors.Is(err, ErrRoomNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Succès - on retourne la room mise à jour
	room, err := s.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
