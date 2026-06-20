package main

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
	"net/http"
	"strings"
)

func TestHandleGetRooms(t *testing.T) {
    store := &RoomStore{}
    store.Add(Room{ID: "101", Name: "Suite", Price: 150})

    req := httptest.NewRequest("GET", "/rooms", nil)
    rec := httptest.NewRecorder()

    store.handleGetRooms(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", rec.Code)
    }

    var rooms []Room
    json.NewDecoder(rec.Body).Decode(&rooms)
    if len(rooms) != 1 {
        t.Errorf("expected 1 room, got %d", len(rooms))
    }
}

func TestHandleAddRoom(t *testing.T) {
	store := &RoomStore{}

	body := strings.NewReader(`{"id":"999", "name":"NewRoom", "price":120}`)
	req := httptest.NewRequest("POST", "/rooms", body)
	rec := httptest.NewRecorder()

	store.handleAddRoom(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}

	if len(store.GetAll()) != 1 {
		t.Errorf("expected 1 room in store, got %d", len(store.GetAll()))
	}
	
	var created Room
	json.NewDecoder(rec.Body).Decode(&created)

	if created.ID != "999" {
		t.Errorf("expected ID 999, got %s", created.ID)
	}
}