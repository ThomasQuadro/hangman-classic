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

func crypt() []string {
	for i := 0; i < len(motjeu); i++ {
		crypto = append(crypto, "_")
	}
	return crypto
}

func main() {
	fmt.Println(mothasard())
	fmt.Println(crypt())
}
