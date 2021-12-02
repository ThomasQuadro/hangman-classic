package main

import "fmt"

func interaction() {
	fmt.Println(string(crypt))       //bon
	fmt.Println("Entrez une Lettre") //bon
	fmt.Scanf("%s", &letter)         //bon
	letter = ToUpper(letter)         //bon
	var l = []rune(letter)
	if verifyIsLetter() { //bon
		fmt.Println("Vous avez choisis la lettre", string(l[0]))
		verifgoodletter() //bon
		position()
		Compteur()
		verifycompleteword()
	} else {
		fmt.Println("Désolé, lettre indisponible")
		interaction()
	}
}
