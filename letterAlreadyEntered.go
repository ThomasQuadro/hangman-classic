package main

import "fmt"

func letterAlreadyEntered() bool {
	var motcompt int
	var r = []rune(letter)
	for i := 0; i < len(letterenter); i++ {
		if r[0] == letterenter[i] {
			motcompt += 1
		} else {
			motcompt += 0
		}
	}
	if motcompt != 0 {
		fmt.Println("La lettre", letter, "à déjà était entrée !")
		return true
	} else {
		letterenter = append(letterenter, r[0])
		return false
	}
}
