package main

import (
	"math/rand"
)

//Fonction qui prend une lettre du mot sélectionné aléatoirement pour commencer le hangman
func motvisiblefirst() string {
	var n = (len(mot) / 2) - 1
	for i := 0; i < n; i++ {
		var alea = rand.Intn(len(mot))
		crypt[alea] = mr[alea]
	}
	return string(crypt)
}
