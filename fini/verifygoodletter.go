package main

//fonction qui compte les erreurs
func verifgoodletter() string {
	var r = []rune(letter)
	var c = 0
	for i := 0; i < len(mr); i++ {
		if mr[i] == r[0] {
			c += 1
			crypt[i] = rune(mr[i])
		} else {
			c += 0
		}
	}
	if c < 1 {
		error--
	} else {
		error += 0
	}
	return string(crypt)
}
