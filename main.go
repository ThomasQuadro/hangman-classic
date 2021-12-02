package main

import (
	"fmt"
	"os"
)

var letter string
var crypt = []rune{'_', '_'}

func interaction() {
	fmt.Println("Entrez une Lettre")
	fmt.Scanf("%s", &letter)
	verifyIsLetter()
	verifycompleteword()
}

func verifyIsLetter() {
	var l = []rune(letter)
	if l[0] <= rune('Z') && l[0] >= rune('A') && len(l) == 1 {
		fmt.Println("Vous avez choisis la lettre", string(l[0]))
	} else {
		fmt.Println("Désolé, lettre indisponible")
	}
}
func verifycompleteword() {
	c := 0
	for i := 0; i < len(crypt); i++ {
		if crypt[i] != '_' {
			c += 1
		} else {
			c += 0
		}
	}
	if c == len(crypt) {
		fmt.Println("GGWP")
		os.Exit(0)
	}
}

func main() {
	interaction()
}
