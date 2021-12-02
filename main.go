package main

import "fmt"

var letter string

func interaction() {
	fmt.Println("Entrez une Lettre")
	fmt.Scanf("%s", &letter)
	verifyIsLetter()
}

func verifyIsLetter() {
	var l = []rune(letter)
	if l[0] <= rune('Z') && l[0] >= rune('A') && len(l) == 1 {
		fmt.Println("Vous avez choisis la lettre", string(l[0]), "Bon choix")
	} else {
		fmt.Println("Désolé, lettre indisponible")
	}
}

func main() {
	interaction()
}
