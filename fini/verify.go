package main

import (
	"fmt"
	"os"
)

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
		fmt.Println("Bravo vous avez trouvé le mot --> ", mot)
		os.Exit(0)
	}
}
