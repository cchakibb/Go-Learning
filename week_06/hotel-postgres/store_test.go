package main

import (
    "database/sql"
    "testing"

    _ "github.com/jackc/pgx/v5/stdlib"
)

// testDB est la connexion partagée entre tous les tests.
var testDB *sql.DB

// TestMain est appelé automatiquement par go test avant tous les tests.
// C'est ici qu'on initialise la connexion et qu'on nettoie après.
func TestMain(m *testing.M) {
    var err error
    testDB, err = sql.Open("pgx",
        "host=localhost port=5433 user=hotel password=secret dbname=hoteldb sslmode=disable")
    if err != nil {
        panic(err)
    }
    if err = testDB.Ping(); err != nil {
        panic("PostgreSQL not reachable: " + err.Error())
    }

    // Lance tous les tests, puis ferme la connexion
    m.Run()
    testDB.Close()
}

// cleanup supprime les rooms de test pour repartir d'un état propre
func cleanup(t *testing.T) {
    t.Helper()
    _, err := testDB.Exec("DELETE FROM rooms WHERE id LIKE 'TEST-%'")
    if err != nil {
        t.Fatalf("cleanup failed: %v", err)
    }
}

func TestAdd(t *testing.T) {
    cleanup(t)
    s := &PostgresStore{db: testDB}

    tests := []struct {
        name    string
        room    Room
        wantErr bool
    }{
        {"valid room", Room{ID: "TEST-01", Name: "Suite", Price: 150}, false},
        {"duplicate ID", Room{ID: "TEST-01", Name: "Other", Price: 200}, true},
        {"empty ID", Room{ID: "", Name: "Suite", Price: 150}, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := s.Add(tt.room)
            if tt.wantErr && err == nil {
                t.Errorf("expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
        })
    }
}

func TestGetByID(t *testing.T) {
    cleanup(t)
    s := &PostgresStore{db: testDB}
    s.Add(Room{ID: "TEST-01", Name: "Suite", Price: 150})

    tests := []struct {
        name    string
        id      string
        wantErr bool
    }{
        {"existing room", "TEST-01", false},
        {"non-existing room", "TEST-99", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            room, err := s.GetByID(tt.id)
            if tt.wantErr {
                if err == nil {
                    t.Errorf("expected error, got none")
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
                if room.ID != tt.id {
                    t.Errorf("got ID %q, want %q", room.ID, tt.id)
                }
            }
        })
    }
}

func TestDelete(t *testing.T) {
    cleanup(t)
    s := &PostgresStore{db: testDB}
    s.Add(Room{ID: "TEST-01", Name: "Suite", Price: 150})

    tests := []struct {
        name    string
        id      string
        wantErr bool
    }{
        {"existing room", "TEST-01", false},
        {"non-existing room", "TEST-99", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := s.Delete(tt.id)
            if tt.wantErr && err == nil {
                t.Errorf("expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
        })
    }
}

func TestUpdatePrice(t *testing.T) {
    cleanup(t)
    s := &PostgresStore{db: testDB}
    s.Add(Room{ID: "TEST-01", Name: "Suite", Price: 150})

    tests := []struct {
        name     string
        id       string
        price    float64
        wantErr  bool
    }{
        {"valid update", "TEST-01", 299.99, false},
        {"non-existing room", "TEST-99", 100, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := s.UpdatePrice(tt.id, tt.price)
            if tt.wantErr && err == nil {
                t.Errorf("expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            // Vérifie que le prix a bien changé en base
            if !tt.wantErr {
                room, _ := s.GetByID(tt.id)
                if room.Price != tt.price {
                    t.Errorf("price not updated: got %.2f, want %.2f",
                        room.Price, tt.price)
                }
            }
        })
    }
}