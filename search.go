package main

import (
	"bufio"
	"log"
	"os"
)

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
