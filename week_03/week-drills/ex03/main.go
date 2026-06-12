package main

import (
	"strings"
	"fmt"
)

type Pipeline struct {
	Name 		string
	Process		func(input string) string
}

func main() {

	pipelines := []Pipeline{
		{Name: "uppercase", Process: func(input string) string {
			return strings.ToUpper(input)
		}}, 
	{Name: "reverse", Process: func(input string) string{
		runes := []rune(input) // converting a string to a slice of runes to make it modiable
		reversedString := []rune{}
		for i:= len(runes) -1; i >= 0; i-- {
			reversedString = append(reversedString, runes[i])
		}
		return string(reversedString)
	}}}

	for _, p := range pipelines {
		fmt.Println(p.Process("hello"))

	}
}
