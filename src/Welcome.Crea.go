package ProjetRed

import (
	"fmt"
	"strings"
)

func (p Player) WelcomeScreen() {
	fmt.Println("Bienvenu dans le Projet Red", "\n", "1: Commencer une nouvelle partie", "\n", "0: Quitter")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		p.CharCrea()
	case "0":
		ClearTerminal()
		fmt.Println("A la prochaine !")
	default:
		ClearTerminal()
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.WelcomeScreen()
	}
}
func isletter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
func keeplettersonly(input1 string) string {
	var result string
	for _, char := range input1 {
		if isletter(char) {
			result += string(char)
		}
	}
	return result
}

func (p *Player) CharCrea() {
	var input1 string
	fmt.Println("Vous allez créer votre propre personnage, pour commencer, entrez un nom (Limité à 15 caractères composé uniquement de lettres)")
	fmt.Scanln(&input1)
	input1 = keeplettersonly(input1)
	if len(input1) > 15 {
		input1 = input1[:15]
	}
	if len(input1) == 0 {
		fmt.Println("Faite un effort, entrez au moin une lettre valide, histoire de dire")
		p.CharCrea()
	} else {
		input1 = strings.ToLower(input1)
		input1 = strings.Title(input1)
		fmt.Println(input1, "sera donc votre nom..")
		var input2 string
		fmt.Println("Maintenant, choississez une race parmis les ces trois choix")
		fmt.Println("\n", "1: Humain", "\n", "2: Elfe", "\n", "3: Nain")
		fmt.Println("Les Humains commences avec 100 PV max, 50 PM max", "\n", "Les Elfes commences avec 80 PV max, 75 PM max", "\n", "Les Nains commences avec 120 PV max, 25 PM max")
		fmt.Scanln(&input2)
		switch input2 {
		case "1":
			ClearTerminal()
			input2 = "Humain"
		case "2":
			ClearTerminal()
			input2 = "Elfe"
		case "3":
			ClearTerminal()
			input2 = "Nain"
		default:
			ClearTerminal()
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.CharCrea()
		}
		Initplayer(input1, input2)
	}
}
