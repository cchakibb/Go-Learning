// 1. Déclare une interface `Describer` avec deux méthodes :
//    - Describe() string
//    - Category() string

// 2. Crée deux structs : `Hotel` (Name string, Stars int)
//    et `Hostel` (Name string, DormSize int)

// 3. Implémente Describer sur les deux :
//    - Hotel.Describe()  → "Hotel [Name], [Stars] stars"
//    - Hotel.Category()  → "luxury" si Stars >= 4, sinon "standard"
//    - Hostel.Describe() → "Hostel [Name], [DormSize] beds"
//    - Hostel.Category() → "budget"

// 4. Écris une fonction `printInfo(d Describer)` qui affiche 
//    les deux infos

// 5. Dans main, teste avec un Hotel 5 étoiles, 
//    un Hotel 3 étoiles, et un Hostel

package main

import "fmt"

// interface
type Describer interface {
	Describe()		string
	Category()		string
}

// structs
type Hotel struct {
	Name			string
	Stars			int
}
type Hostel struct {
	Name			string
	DormSize		int
}
func (hot Hotel) Describe() string {
	return fmt.Sprintf("Hotel %s, %d stars", hot.Name, hot.Stars)
}
func (hot Hotel) Category() string {
	if hot.Stars >= 4 {
		return "luxury"
	}
	return "standard"
}
func (host Hostel) Describe() string {
	return fmt.Sprintf("Hostel %s, %d beds", host.Name, host.DormSize)
}
func (host Hostel) Category() string {
	return "budget"
}
func PrintInfo(d Describer) {
	fmt.Println(d.Describe(), d.Category())
}

func main() {

	hotel := Hotel{Name: "Saint Aignan", Stars: 5}
	hotel2 := Hotel{Name: "Best Western", Stars: 3}
	hostel := Hostel{Name: "L'Auberge", DormSize: 325}

	PrintInfo(hotel)
	PrintInfo(hotel2)
	PrintInfo(hostel)
}