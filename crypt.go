package main

//Fonction qui dresse des lignes a la place des lettres
func crypto() string {
	for i := 0; i < len(mot); i++ {
		crypt = append(crypt, '_')
	}
	motvisiblefirst()
	return string(crypt)
}
