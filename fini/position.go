package main

import "fmt"

//position du hangman selon le nombre d'erreur
func position() {
	var pos string
	for i := 0; i < 8; i++ {
		if error == 9 {
			pos += search()[i]
		}
		if error == 8 {
			pos += search()[8+i]
		}
		if error == 7 {
			pos += search()[16+i]
		}
		if error == 6 {
			pos += search()[24+i]
		}
		if error == 5 {
			pos += search()[32+i]
		}
		if error == 4 {
			pos += search()[40+i]
		}
		if error == 3 {
			pos += search()[48+i]
		}
		if error == 2 {
			pos += search()[56+i]
		}
		if error == 1 {
			pos += search()[64+i]
		}
		if error == 0 {
			pos += search()[72+i]
		}
	}
	fmt.Println(pos)
}
