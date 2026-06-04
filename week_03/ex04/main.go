package main

import (
    "fmt"
    "sync"
)

var mu sync.Mutex
var counter int
var wg sync.WaitGroup

func increment() {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
		mu.Lock()
        counter++
		mu.Unlock()
    }
}

func main() {
    wg.Add(5)
    for i := 0; i < 5; i++ {
        go increment()
    }
    wg.Wait()
    fmt.Println("counter:", counter)
}
