package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var motjeu string
var crypto []string

//Fonction qui prend un mot aléatoirement dans la liste de mots .txt
func mothasard() string {
	file, err := ioutil.ReadFile("word/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	answers := string(file)
	words := strings.Split(answers, "\n")
	rand.Seed(time.Now().UnixNano())
	motjeu = words[rand.Intn(len(words))]
	return motjeu
}

//Fonction qui dresse des lignes a la place des lettres
func crypt() string {
	var line string
	for i := 0; i < len(motjeu); i++ {
		crypto = append(crypto, "_")
		line = line + crypto[i] + " "
	}
	return line
}

//Fonction qui prend une lettre du mot sélectionné aléatoirement pour commencer le hangman
func motvisiblefirst() {
	var alea = rand.Intn(len(motjeu))
	// fmt.Println(alea)
	// fmt.Println(string(motjeu[alea]))
	crypto[alea] = string(motjeu[alea])
	fmt.Println(crypt())

}

//Fonction print
func main() {
	fmt.Println(mothasard())
	fmt.Println(crypt())
	motvisiblefirst()
}
