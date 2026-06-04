package main

import (
	"fmt"
	"math/rand"
)

func main() {

	source1 := Source{
		Name: "DB",
		Check: func(roomID string) (bool, error) {
			return rand.Intn(2) == 1, nil
		},
	}

	source2 := Source{
		Name: "Cache",
		Check: func(roomID string) (bool, error) {
			return rand.Intn(2) == 1, nil
		},
	}

	source3 := Source{
		Name: "API",
		Check: func(roomID string) (bool, error) {
			return rand.Intn(2) == 1, nil
		},
	}

	sources := []Source{source1, source2, source3}

	ch := make(chan Result)

	fmt.Printf("Checking availability for room 101...\n\n")
	for _, s := range sources{
		go checkSource(s, "101", ch)
	}

	counter := 0
	for i := 0; i < len(sources); i++ {
		res := <- ch

		if res.Available {
			counter++
		}
		if res.Available {
			fmt.Printf("[%s] room 101: available (%v)\n", res.Source, res.Duration)
		} else {
			fmt.Printf("[%s] room 101: unavailable (%v)\n", res.Source, res.Duration)
		}
	}
	fmt.Printf("\nResults: %d/%d sources say available\n", counter, len(sources))
}