package main

import (
	"testing"
	"errors"
)

func TestCheckSourceAvailable(t * testing.T){
	ch := make(chan Result)

	source := Source{
		Name: "TestSource",
		Check: func(roomID string) (bool, error) {
			return true, nil
		},
	}
	go checkSource(source, "101", ch)
	res := <- ch
	if !res.Available {
		t.Errorf("Available: %t but expected true", res.Available)
	}
	if res.Source != "TestSource" {
		t.Errorf("Source name: %s but expected TestSource", res.Source)
	}
	if res.err != nil {
		t.Errorf("Expected err = nil but we have err = %s", res.err)
	}

}

func TestCheckSourceError(t * testing.T){
	ch := make(chan Result)

	source := Source{
		Name: "TestSourceError",
		Check: func(roomID string) (bool, error) {
			return false, errors.New("Source down")
		},
	}
	go checkSource(source, "101", ch)
	res := <- ch

	if res.Available {
		t.Errorf("Available: %t but expected %t", res.Available, !res.Available)
	}
	if res.err == nil {
		t.Errorf("We should have an error here but we don't: we have res.err == nil")
	}
}