package main

import (
	"fmt"
	"os"
)

var mot = "LIT"
var mr = []rune(mot)

var letter = "I"
var r = []rune(letter)

var crypt = []rune{'_', '_', '_'}

var error int = 10

func Compteur() {
	if error == 0 {
		fmt.Println("Vous avez perdu")
		os.Exit(0)
	}
}

func verifgoodletter() {
	for i := 0; i < len(mr); i++ {
		if mr[i] == r[0] {
			crypt[i] = r[0]
		} else {
			error--
		}
	}
}
