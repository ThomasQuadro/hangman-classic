package main

import (
	"fmt"
	"os"
)

func verifyWord() {
	if mot == letter {
		fmt.Println("Bravo vous avez trouvé le mot --> ", mot)
		os.Exit(0)
	} else {
		fmt.Println(affichage())
		error -= 2
		fmt.Println("Triste mais ce n'est pas le bon mot XD\nVous avez perdu 2 erreurs\nIl vous reste", error, "erreurs")
		position()
		interaction()
	}
}
