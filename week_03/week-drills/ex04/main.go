// 1. Déclare une interface `Formatter` avec
//    une méthode Format(s string) string

// 2. Déclare une struct `Printer` avec :
//    - un champ Transform func(string) string
//    - une méthode Format(s string) string
//      qui appelle p.Transform(s) et retourne le résultat

// 3. Dans main :
//    - Crée une instance Printer avec Transform = strings.ToUpper
//    - Appelle p.Format("hello") directement
//    - Passe p à une fonction printFormatted(f Formatter, s string)
//      qui appelle f.Format(s) et affiche le résultat

// 4. Lance et affiche le résultat

//L'objectif : Printer satisfait Formatter via sa méthode,
// mais utilise son champ fonction en interne. Les deux mécanismes dans le même type.

package main

import ("strings"
		"fmt")


type Formatter interface {
    Format(s string) string // A : contrat — "quiconque veut être un Formatter doit avoir cette méthode"
}

type Printer struct {
    Transform func(string) string // B : donnée stockée — la logique concrète, variable par instance
}

func (p Printer) Format(s string) string { // C : Printer honore le contrat de Formatter
    return p.Transform(s) // D : la méthode délègue à la donnée interne
}

func printFormatted(f Formatter, s string){
	fmt.Println(f.Format(s))
}

func main(){
	p := Printer{Transform: strings.ToUpper}
	fmt.Println(p.Format("hello"))
	printFormatted(p, "bye")
}