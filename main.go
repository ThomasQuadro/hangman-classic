package main

import "fmt"

func main() {
	var i rune
RETOUR:
	fmt.Println("Entrez une Lettre")
	fmt.Scan(&i)
	// si c'est vrais

	// Sinon
	//fmt.Println("Vous avez choisis la lettre", i, "Mauvais choix")

	if i > rune('Z') || i < rune('A') {
		fmt.Println("Désolé, lettre indisponnible")
		goto RETOUR
	} else if i < rune('Z') || i > rune('A') {
		fmt.Println("Vous avez choisis la lettre", string(i), "Bon choix")
	}

}
