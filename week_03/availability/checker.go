package main

import (
	"time"
	"math/rand"
)

type Source struct {
	Name 		string
	Check		func(roomID string) (bool, error)
}

type Result struct {
	Source 		string
	Available 	bool
	Duration 	time.Duration
	err 		error
}

func checkSource(source Source, roomID string, ch chan Result){

	duration := time.Duration(rand.Intn(498) + 100) * time.Millisecond
	time.Sleep(duration)
	
	isAvailable, errSource := source.Check(roomID)
	result := Result{Source: source.Name, Available: isAvailable, Duration: duration, err: errSource}
	ch <- result
}