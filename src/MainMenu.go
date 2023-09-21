package ProjetRed

import "fmt"

func (p *Player) MainMenu() {
	var input string
	fmt.Println("----------------------------")
	fmt.Println("Menu principal:")
	fmt.Println("------")
	fmt.Println("Information personnage: 1")
	fmt.Println("Inventaire: 2")
	fmt.Println("Marchand: 3")
	fmt.Println("Forgeron: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("----------------------------")
		p.DisplayPlayerInfo()
	case "2":
		fmt.Println("----------------------------")
		p.AccessInventory()
	case "3":
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le marchand !", "\n", "Je vous propose quelques objet pour votre aventure, je peux également vous racheter vos objet")
		fmt.Println("------")
		p.Dealer()
	case "4":
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le forgeron !", "\n", "Contre une légère commission, ma forge est à votre disposition.")
		fmt.Println("------")
		p.BlackSmith()
	case "0":
		fmt.Println("----------------------------")
		fmt.Println("A plus !")
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.MainMenu()
	}
}
