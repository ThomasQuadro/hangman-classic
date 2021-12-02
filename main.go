package main

import "fmt"

var letter = "I"
var l = []rune(letter)

func verifyIsLetter() {

	var i rune
RETOUR:
	fmt.Println("Entrez une Lettre")
	fmt.Scan(&i)
	if i > rune('Z') || i < rune('A') || i > rune('z') || i < rune('a') {
		fmt.Println("Désolé, lettre indisponnible")
		goto RETOUR

	} else {
		fmt.Println("Vous avez choisis la lettre", string(i), "Bon choix")
	}
}

func main() {
	fmt.Println(l)
	// fmt.Println(verifyIsLetter())
}
