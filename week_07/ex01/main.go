package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, resultChan chan <- string) {
	// simulate a 3 second process
	select {
	case <- time.After(1 * time.Second):
		resultChan <- "work done"
	case <- ctx.Done():
		return // we stop here, no sending stuck on resultChan
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	resultChan := make(chan string)
	go doWork(ctx, resultChan)

	select {
	case res := <- resultChan:
		fmt.Println("Received result :", res)
	case <- ctx.Done():
		fmt.Println("Timeout ! Error :", ctx.Err())
	}
}