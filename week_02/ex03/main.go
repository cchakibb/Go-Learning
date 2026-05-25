// Reprends le code de l'exercice 2.
// 1. Crée un slice []Describer avec 2 Hotels et 2 Hostels mélangés
// 2. Itère avec range et appelle PrintInfo sur chaque élément
// 3. Compte combien sont "luxury" et affiche le total à la fin

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

	fmt.Println()

	properties := []Describer {
		Hotel{Name: "Hilton", Stars: 5},
		Hostel{Name: "BackpackInn", DormSize: 20},
		Hotel{Name: "Ibis", Stars: 3},
		Hostel{Name: "YouthInn", DormSize: 84},
	}
	luxuryCounter := 0

	for _, d := range properties {
		PrintInfo(d)
		if d.Category() == "luxury" {
			luxuryCounter++
		}
	}
	fmt.Println("Total luxury:", luxuryCounter)
}