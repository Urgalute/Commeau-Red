package ProjetRed

import (
	"fmt"
)

var input int

func (p *Player) Dealer() {
	fmt.Println("Page 1: Utilitaires")
	fmt.Println("------")
	fmt.Println("Acheter une potion de PV, 3 Pièces d'or: 1")
	fmt.Println("Potion de poison, 6 Pièces d'or: 2")
	fmt.Println("Livre de sort: Boule de feu, 25 Pièces d'or : 3")
	fmt.Println("Page 2: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddInventory("Potion de PV", 1, 3)
		p.Dealer()
	case 2:
		p.AddInventory("Potion de poison", 1, 6)
		p.Dealer()
	case 3:
		p.AddInventory("Livre de sort: Boule de feu", 1, 25)
		p.Dealer()
	case 4:
		p.Dealer2()
	case 0:
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer()
	}
}
func (p *Player) Dealer2() {
	fmt.Println("----------------------------")
	fmt.Println("Page 2: Composants")
	fmt.Println("------")
	fmt.Println("Fourrure de loup, 4 Pièces d'or: 1")
	fmt.Println("Peau de troll, 7 Pièces d'or: 2")
	fmt.Println("Cuir de sanglier, 3 Pièces d'or: 3")
	fmt.Println("Plume de corbeau, 1 Pièces d'or: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddInventory("Fourrure de loup", 1, 4)
		p.Dealer2()
	case 2:
		p.AddInventory("Peau de troll", 1, 7)
		p.Dealer2()
	case 3:
		p.AddInventory("Cuir de sanglier", 1, 3)
		p.Dealer2()
	case 4:
		p.AddInventory("Plume de corbeau", 1, 1)
		p.Dealer2()
	case 0:
		p.Dealer()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer2()
	}
}
