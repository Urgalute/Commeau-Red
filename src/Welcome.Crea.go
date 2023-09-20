package ProjetRed

import (
	"fmt"
)

func (p Player) WelcomeScreen() {
	fmt.Println("Bienvenu dans le Projet Red", "\n", "Commencer une nouvelle partie: 1", "\n", "Quitter: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.CharCrea()
	case 0:
		fmt.Println("A la prochaine !")
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.WelcomeScreen()
	}
}
func (p *Player) CharCrea() {
	var input1 string
	var input2 string
	var input3 string
	fmt.Println("Vous allez créer votre propre personnage, pour commencer, entrez un nom")
	fmt.Scanln(&input1)
	i := input1
	fmt.Println(i, "sera donc votre nom..")
	fmt.Println("Maintenant, choississez une race parmis les ces trois choix")
	fmt.Scanln(&input2)
	fmt.Scanln(&input3)
	Initplayer(input1, input2, input3)
}
