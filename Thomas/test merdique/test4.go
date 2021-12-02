package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var filePath = "words.txt"
var words = []string{}

func ReadWordsFromFile(filePath string) []string {
	b, err := ioutil.ReadFile(filePath) // read words from file
	if err != nil {
		log.Fatal(err)
	}

	str := string(b) // convert content to a 'string'
	words = strings.Split(str, "\n")
	return words
}

func main() {
	fmt.Println(ReadWordsFromFile(filePath))
	fmt.Println(words[1])
}
