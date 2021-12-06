package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

//Fonction qui prend un mot aléatoirement dans la liste de mots .txt
func mothasard() string {
	file, err := ioutil.ReadFile("word/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	answers := string(file)
	words := strings.Split(answers, "\n")
	rand.Seed(time.Now().UnixNano())
	mot = ToUpper(words[rand.Intn(len(words))])
	return mot
}
