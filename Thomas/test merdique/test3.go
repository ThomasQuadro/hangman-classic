package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func random_word(numletters int, gamemode int) string {
	switch gamemode {
	case 1:
		var dataletters []byte
		var err error
		if numletters == 4 {
			dataletters, err = ioutil.ReadFile("word/words2.txt")
		} else if numletters == 5 {
			dataletters, err = ioutil.ReadFile("word/words2.txt")
		} else if numletters >= 6 {
			dataletters, err = ioutil.ReadFile("word/words2.txt")
		}

		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword

	case 2:
		var dataletters []byte
		var err error
		if numletters == 4 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 5 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 6 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 7 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 8 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 9 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 10 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 11 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 12 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 13 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 14 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		} else if numletters == 15 {
			dataletters, err = ioutil.ReadFile("word/words3.txt")
		}

		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword

	}

	return "omgthisisabugyoushouldntseethisever"
}

func main() {
	fmt.Println(random_word(2, 4))
}
