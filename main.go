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

	if i > rune(91) || i < rune(64) {
		fmt.Println("Désolé, lette indisponnible")
		goto RETOUR
	} else {
		fmt.Println("Vous avez choisis la lettre", string(i), "Bon choix")
	}

}
