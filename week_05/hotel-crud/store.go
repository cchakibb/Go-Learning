package main

import (
	"errors"
	"sync"
)

type Room struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type RoomStore struct {
	mu    sync.Mutex
	rooms []Room
}

func (s *RoomStore) Add(room Room) {
	// lock, append, unlock
	s.mu.Lock()
	s.rooms = append(s.rooms, room)
	s.mu.Unlock()

}

func (s *RoomStore) GetAll() []Room {
	// lock, copy, unlock, return copy
	s.mu.Lock()
	newRooms := append([]Room{}, s.rooms...)
	s.mu.Unlock()
	return newRooms
}

func (s *RoomStore) GetByID(id string) (Room, error) {
	// searches a room by ID. Returns error if not found
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, room := range s.rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return Room{}, errors.New("room not found")
}

func (s *RoomStore) Delete(id string) (Room, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, room := range s.rooms {
		if room.ID == id {
			s.rooms = append(s.rooms[:i], s.rooms[i+1:]...)
			return room, nil
		}
	}
	return Room{}, errors.New("room not found")
}
