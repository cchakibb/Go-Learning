package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	fmt.Println("Hotel CLI. Commands: add <roomID> <guestName> <nights>, list, price <roomID> <basePrice>, cancel <roomID>, quit")

	h := Hotel{}
	p := StandardPricing{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		usrInput := strings.SplitN(line, " ", 4) // add <roomID> <guestName> <nights>
		command := usrInput[0]

		switch command {
		case "add":
			if len(usrInput) < 4 {
				fmt.Println("usage: add <roomID> <guestName> <nights>")
			} else {
				nights, err := strconv.Atoi(usrInput[3])
				if err == nil {
					h.Add(usrInput[1], usrInput[2], nights)
				} else {
					fmt.Println(err)
				}
			}
		case "list":
			h.List()
		case "price":
			if len(h.Bookings) >= 1 {
				b, err := h.Find(usrInput[1])
				if err != nil {
					fmt.Println(err)
				} else {
					basePrice, errBasePrice := strconv.ParseFloat(usrInput[2], 64)
					if errBasePrice != nil {
						fmt.Println(errBasePrice)
					} else {
						result := p.Calculate(basePrice, b.Nights)
						fmt.Printf("Price for room %s: %.2f\n", b.RoomID, result)
					}
				}
			}
		case "cancel":
			err := h.Cancel(usrInput[1])
			if err != nil {
				fmt.Println(err)
			}
		case "quit":
			return
		default:
			fmt.Println("This command does not exist. Please use: add <roomID> <guestName> <nights>, list, price <roomID> <basePrice>, cancel <roomID>, quit")
		}
	}
	if err := scanner.Err() ; err != nil {
		fmt.Println("input error: ", err)
	}
}