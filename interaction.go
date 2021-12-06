package main

import (
	"fmt"
)

func interaction() {
	for error >= 0 {
		fmt.Println("Entrez une Lettre") //bon
		fmt.Scanf("%s", &letter)         //bon
		letter = ToUpper(letter)         //bon
		if verifyIsLetter() {            //bon
			if letterAlreadyEntered() == true {
				interaction()
			}
			verifgoodletter()        //bon
			fmt.Println(affichage()) //bon
			position()
			Compteur()
			verifycompleteword()
		} else {
			enterWord()
			verifyWord()
		}
	}
}
