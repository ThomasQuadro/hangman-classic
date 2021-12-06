package main

func enterWord() bool {
	var l = []rune(letter)
	var compt int
	for i := 0; i < len(letter); i++ {
		if l[i] <= rune('Z') && l[i] >= rune('A') && len(l) == len(mot) {
			compt += 1
		} else {
			compt += 0
		}
	}
	if compt == len(letter) {
		return true
	} else {
		return false
	}
}
