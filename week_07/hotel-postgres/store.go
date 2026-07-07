package main

import (
	"database/sql"
	"errors"
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var ErrRoomNotFound = errors.New("room not found")

type Room struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) Add(room Room) error {
	if room.ID == "" {
        return errors.New("room ID cannot be empty")
    }
	_, err := s.db.Exec(
		"INSERT INTO rooms (id, name, price) VALUES ($1, $2, $3)",
		room.ID, room.Name, room.Price,
	)
	return err // nil if success, otherwise err
}

func (s *PostgresStore) GetAll(ctx context.Context) ([]Room, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, price FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var r Room
		if err := rows.Scan(&r.ID, &r.Name, &r.Price); err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, nil
}

func (s *PostgresStore) GetByID(id string) (Room, error) {
	var room Room
	err := s.db.QueryRow(
		"SELECT id, name, price FROM rooms WHERE id = $1", id,
	).Scan(&room.ID, &room.Name, &room.Price)
	if errors.Is(err, sql.ErrNoRows) {
		return Room{}, ErrRoomNotFound
	}
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

func (s *PostgresStore) Delete(id string) error {
	res, err := s.db.Exec("DELETE FROM rooms WHERE id = $1", id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrRoomNotFound
	}
	return nil
}

func (s *PostgresStore) UpdatePrice(id string, newPrice float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM rooms WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return ErrRoomNotFound
	}

	// Met à jour le prix
	_, err = tx.Exec("UPDATE rooms SET price = $1 WHERE id = $2", newPrice, id)
	if err != nil {
		return err
	}
	return tx.Commit() // Valide les deux opérations ensemble (tx.QueryRow + tx.Exec)
}
