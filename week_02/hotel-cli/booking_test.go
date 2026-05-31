package main

import "testing"

func TestAdd(t *testing.T) {
	h := Hotel{}
	h.Add("101", "Chakib", 3)
	h.Add("102", "Naila", 2)

	if len(h.Bookings) != 2 {
		t.Errorf("expected 2 bookings, got %d", len(h.Bookings))
	}
	if h.Bookings[0].RoomID != "101" {
		t.Errorf("expected room 101, got %s", h.Bookings[0].RoomID)
	}
	if h.Bookings[0].Nights != 3 {
		t.Errorf("expected 3 nights, got %d", h.Bookings[0].Nights)
	}
	if h.Bookings[0].GuestName != "Chakib" {
		t.Errorf("expected guest Chakib, got %s", h.Bookings[0].GuestName)
	}

}

func TestCancel(t *testing.T) {
	h := Hotel{}
	// I create a hotel and add reservations, otherwise there is nothing to cancel
	h.Add("101", "Chakib", 5)
	h.Add("102", "Seddik", 15)
	h.Add("103", "Ghenima", 15)

	tests := []struct { // I'm not sure what I'm doing...
		roomID		string
		wantErr		bool
	}{
		{"101", false},
		{"102", false},
		{"999", true},
	}

	for _, tt := range tests {
		lenBookings := len(h.Bookings)
		if tt.wantErr {
			errErr := h.Cancel(tt.roomID)
			if errErr == nil {
				t.Errorf("We should have had an error here, but we didn't")
			}
		} else {
			errErr := h.Cancel(tt.roomID)
			if lenBookings == len(h.Bookings) {
				t.Errorf("We had %d bookings before cancellation. After cancellation we should have %d bookings, but got %d", lenBookings, lenBookings - 1, len(h.Bookings))
			}
			if errErr != nil {
				t.Errorf("%s", errErr)
			}
		}
	}
}

func TestList(t *testing.T) {
    h := Hotel{}
    h.List() // shoudl not panic on empty list
    h.Add("101", "Chakib", 5)
    h.Add("102", "Naila", 3)
    h.List() // should not panic on non empty list
}