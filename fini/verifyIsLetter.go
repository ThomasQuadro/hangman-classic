package main

func verifyIsLetter() bool {
	var l = []rune(letter)
	if l[0] <= rune('Z') && l[0] >= rune('A') && len(l) == 1 {
		return true
	} else {
		return false
	}
}
