package main

import "fmt"

// func swap(x *int, y *int) {
// 	temp := *x
// 	*x = *y
// 	*y = temp
// }

// func main() {
// 	var a int = 4
// 	var b int = 2
// 	fmt.Println(a, b)
// 	swap(&a, &b)
// 	fmt.Println(a, b)
// }

func swap(x, y int) (int, int) {
    return y, x
}

func main() {
    a, b := 4, 2
    a, b = swap(a, b)
    fmt.Println(a, b)
}