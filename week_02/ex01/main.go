// 1. Déclare une interface `Shape` avec une méthode `Area() float64`

// 2. Crée deux structs : `Rectangle` (Width, Height float64) 
//    et `Circle` (Radius float64)

// 3. Implémente `Area()` sur chacune
//    - Rectangle : Width * Height
//    - Circle : 3.14 * Radius * Radius

// 4. Écris une fonction `printArea(s Shape)` qui affiche l'aire

// 5. Dans main, crée un Rectangle et un Circle, 
//    passe-les tous les deux à printArea


package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {

	Width, Height float64
}

type Circle struct {

	Radius float64
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {

	rect := Rectangle{Width: 12.30, Height: 6.30}
	circ := Circle{Radius: 4.30}

	printArea(rect)
	printArea(circ)

}