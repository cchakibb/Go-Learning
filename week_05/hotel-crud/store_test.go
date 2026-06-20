package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := &RoomStore{}
	s.Add(Room{ID: "101", Name: "OceanView", Price: 105.45})

	if len(s.rooms) != 1 {
		t.Errorf("Expected 1, got %d", len(s.rooms))
	}
}

func TestGetByID(t *testing.T) {
	s := &RoomStore{}
	s.Add(Room{ID: "101", Name: "OceanView", Price: 105.45})
	s.Add(Room{ID: "102", Name: "Suite", Price: 155.75})
	s.Add(Room{ID: "315", Name: "Standard", Price: 99.0})

	room, err := s.GetByID("102")
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}
	if room.ID != "102" {
		t.Errorf("Expected room ID: 102 but got %s", room.ID)
	}
	if room.Name != "Suite" {
		t.Errorf("Expected room name Suite but got %s", room.Name)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	s := &RoomStore{}
	s.Add(Room{ID: "101", Name: "PoolSide"})

	_, err := s.GetByID("103")
	if err == nil {
		t.Errorf("Expected an error, resulted with no error !")
	}
}

func TestDelete(t *testing.T) {
	s := &RoomStore{}
	s.Add(Room{ID: "101", Name: "OceanView", Price: 105.45})
	s.Add(Room{ID: "102", Name: "OceanView", Price: 105.45})

	_, err := s.Delete("101")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(s.rooms) != 1 {
		t.Errorf("After deletion, expected 1 room left, but got %d", len(s.rooms))
	}
}
