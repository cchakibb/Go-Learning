package main

import "fmt"

func clean(s string) string {

	runes := []rune(s)
	result := []rune{}

	for _, char := range runes {
		if char < 48 || char > 57 {
			result = append(result, char)
		}
	}
	return string(result)
}


func main() {
	s1 := "h3ll0 w0rld"
	s2 := "Yetna7aw Ga3"
	s3 := "lool 55555 jejeje"

	fmt.Println(clean(s1))
	fmt.Println(clean(s2))
	fmt.Println(clean(s3))
}