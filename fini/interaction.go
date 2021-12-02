package main

import "fmt"

func interaction() {
	fmt.Println(string(crypt))       //bon
	fmt.Println("Entrez une Lettre") //bon
	fmt.Scanf("%s", &letter)         //bon
	letter = ToUpper(letter)         //bon
	verifyIsLetter()                 //bon
	verifgoodletter()                //bon
	position()
	Compteur()
	verifycompleteword()
	interaction()
}
