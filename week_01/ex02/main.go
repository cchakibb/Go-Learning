package main

import "fmt"

func main() {

	var temperature int = 6

	if temperature < 10 {
		fmt.Println("cold")
	} else if temperature >= 10 && temperature <= 25 {
		fmt.Println("ok")
	} else {
		fmt.Println("hot")
	}

	for i := 1; i <= temperature; i++ {
		fmt.Println(i)
	}

	switch temperature / 10 {
	case 0:
		fmt.Println("freezeing")
	case 1:
		fmt.Println("comfortable")
	case 2:
		fmt.Println("warm")
	default:
		fmt.Println("extreme")
	}
}
