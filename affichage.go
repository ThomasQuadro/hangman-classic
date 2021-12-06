package main

func affichage() string {
	var tabaffiche []rune
	for i := 0; i < len(crypt); i++ {
		tabaffiche = append(tabaffiche, crypt[i])
		tabaffiche = append(tabaffiche, ' ')
	}
	return string(tabaffiche)
}
