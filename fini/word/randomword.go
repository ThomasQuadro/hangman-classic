package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var motjeu string
var crypto []string

// var lettrevalides int

func search() []string {
	readFile, err := os.Open("position/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, "\n"+fileScanner.Text())
	}
	readFile.Close()
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	return lines
}

//Fonction qui transforme en Majuscule
func ToUpper(s string) string {
	chaine := []rune(s)
	for i := range chaine {
		if chaine[i] >= 'a' && chaine[i] <= 'z' {
			chaine[i] = chaine[i] - 32
		}
	}
	return string(chaine)
}

//Fonction qui prend un mot aléatoirement dans la liste de mots .txt
func mothasard() string {
	file, err := ioutil.ReadFile("word/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	answers := string(file)
	words := strings.Split(answers, "\n")
	rand.Seed(time.Now().UnixNano())
	motjeu = ToUpper(words[rand.Intn(len(words))])
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
	var n = len(motjeu)/2 - 1
	for i := 0; i < n; i++ {
		var alea = rand.Intn(len(motjeu))
		// fmt.Println(alea)
		// fmt.Println(string(motjeu[alea]))
		crypto[alea] = string(motjeu[alea])
	}
	crypt()
}

//Fonction print
func main() {
	mothasard()
	fmt.Println(motjeu)
	crypt()
	motvisiblefirst()
	// println(len(motjeu))
	interaction()
}

//Fonction qui verifie si c'est une lettre
func verifylongueur(str string) bool {
	if len(str) < 2 && len(str) != 0 {
		return true
	} else {
		return false
	}
}

//fonction qui vérifie si la lettre est dans le mot
func verifyIsLetter(str string) bool {
	var tabrune []rune
	for _, letter := range str {
		tabrune = append(tabrune, letter)
		// println(tabrune)
	}
	for i := 0; i < len(tabrune); i++ {
		if tabrune[i] >= 'A' && tabrune[i] <= 'Z' {
			// println("Vrai")
			return true
		} else {
			// println("faux")
			return false
		}
	}
	return false
}

var f []bool     //tableau qui met false si la lettre n'est pas dans le mot et vraie si elle est dans le mot
var verif []bool //tableau de false = à la longueur du mot
var nbr int      //compteur qui compare le nombre de false entre f et verif, si != alors lettre est dans le mot ( on compare len(mot) avec nbr)

//fonction qui verifie si la lettre est dans le mot
func verifyIsGoodLetter(str string) {
	good_Word()
	if errore > 0 {
		//tableau verif
		for i := 0; i < len(motjeu); i++ {
			verif = append(verif, false)
		}

		//tab f
		for i := 0; i < len(motjeu); i++ {
			if rune(str[0]) != rune(motjeu[i]) {
				f = append(f, false)
			} else {
				f = append(f, true)
			}
		}

		//nbr
		for i := 0; i < len(motjeu); i++ {
			if f[i] == verif[i] {
				nbr += 1
			} else {
				nbr += 0
			}
		}

		//on compare nbr a len(motjeu) pour rajouter la lettre ou non dans le mot avec les tirets
		if nbr != len(motjeu) {
			for i := 0; i < len(motjeu); i++ {
				if rune(str[0]) == rune(motjeu[i]) {
					crypto[i] = string(rune(str[0]))
				}
			}
			fmt.Println("La lettre", str, "est dans le mot")
		} else {
			fmt.Println("La lettre", str, "n'est pas dans le mot")

		}
		// fmt.Println(crypt())
	} else {
		fmt.Println("Vous avez épuisé votre jauge d'erreur !")
		os.Exit(0)
	}

}

//notre lettre que l'on va écrire
var letter string

// fonction interaction homme-machine
func interaction() {
	fmt.Println(crypt())
	fmt.Println("Entrez une lettre :")
	fmt.Scanf("%s", &letter)
	//on convertie i en majuscule
	letter = ToUpper(letter)
	if verifylongueur(letter) && verifyIsLetter(letter) {
		lettresortie()
		verify_error()
		fmt.Println("Vous avez choisie la lettre", letter)
		verifyIsGoodLetter(letter)
		fmt.Println("-----------------------------------")
		interaction()
	} else {
		fmt.Println("Veuillez entrez une lettre !")
	}
}

var lettresort []string

func lettresortie() {
	for i := 0; i < len(lettresort); i++ {
		if letter != lettresort[i] {
			lettresort = append(lettresort, letter)
		} else {
			fmt.Println("Cette lettre est deja sortie")
			fmt.Println("Veuillez rentrez une lettre qui n'est pas sortie !")
			interaction()
		}
	}
	fmt.Println(lettresort)
}

// verifie le nombre d'erreur et met les positionnements du hangman
var errore = 10

func verify_error() {
	var compteur = 0
	for j := 0; j < len(motjeu); j++ {
		if letter[0] == motjeu[j] {
			compteur += 0
		} else {
			compteur += 1
		}
	}
	if compteur == len(motjeu) {
		errore--
		fmt.Println("Oh NON ! Vous avez fait une erreur !")
		fmt.Println("Vous avez le droit à", errore, "erreurs")
	} else {
		fmt.Println("Bien joué vous n'avez pas fait d'erreur !")
		fmt.Println("Vous avez le droit à", errore, "erreurs")
	}
	position()
}

//position du hangman selon le nombre d'erreur
func position() {
	if errore == 9 {
		fmt.Println(search()[0:7])
	}
	if errore == 8 {
		fmt.Println(search()[8:15])
	}
	if errore == 7 {
		fmt.Println(search()[16:23])
	}
	if errore == 6 {
		fmt.Println(search()[24:31])
	}
	if errore == 5 {
		fmt.Println(search()[32:39])
	}
	if errore == 4 {
		fmt.Println(search()[40:47])
	}
	if errore == 3 {
		fmt.Println(search()[48:55])
	}
	if errore == 2 {
		fmt.Println(search()[56:63])
	}
	if errore == 1 {
		fmt.Println(search()[64:71])
	}
	if errore == 0 {
		fmt.Println(search()[72:79])
	}
}

//fonction qui permet de voir si le mot est complet
var motbon = 0

func good_Word() {
	for k := 0; k < len(motjeu); k++ {
		if motjeu[k] == crypt()[k] {
			motbon += 1
		}
	}
	if motbon == len(motjeu) {
		os.Exit(0)
	}
	fmt.Println(motbon)
}
