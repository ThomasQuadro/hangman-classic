package main

import "fmt"

func verifyIsLetter() {
	var l = []rune(letter)
	if l[0] <= rune('Z') && l[0] >= rune('A') && len(l) == 1 {
		fmt.Println("Vous avez choisis la lettre", string(l[0]))
	} else {
		fmt.Println("Désolé, lettre indisponible")
		interaction()
	}
}
