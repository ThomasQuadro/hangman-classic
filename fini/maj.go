package main

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
