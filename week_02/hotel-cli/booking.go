package main

import ("errors"
		"fmt")

type Booking struct {
	RoomID		string
	GuestName	string
	Nights		int
}

type Hotel struct {
	Bookings []Booking
}

// Add a new booking
func (h *Hotel) Add(roomID string, guestName string, nights int) {
	b := Booking{RoomID: roomID, GuestName: guestName, Nights: nights}
	h.Bookings = append(h.Bookings, b)
}

func (h *Hotel) Find(roomID string) (Booking, error) {
	for _, b := range h.Bookings {
		if b.RoomID == roomID {
			return b, nil
		}
	}
	return Booking{}, errors.New("room not found")
}

// Cancel a booking - error if not found
func (h *Hotel) Cancel(roomID string) error {
	
	for i, b := range h.Bookings {	
		if b.RoomID == roomID {
			h.Bookings = append(h.Bookings[:i], h.Bookings[i+1:]...)
			return nil
		}
	}
	return errors.New("Reservation not found")
}

// show all bookings
func (h *Hotel) List() {
	for _, b := range h.Bookings {
		fmt.Printf("Room: %s, Guest: %s, Nights: %d\n", b.RoomID, b.GuestName, b.Nights)
	}
}