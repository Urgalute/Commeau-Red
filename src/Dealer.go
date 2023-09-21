package ProjetRed

import (
	"fmt"
)

var input string

func (p *Player) Dealer() {
	fmt.Println("Page 1: Utilitaires")
	fmt.Println("------")
	fmt.Println("Votre or:", p.money)
	fmt.Println("Acheter une potion de PV, 3 Pièces d'or: 1")
	if p.inventory["Potion de PV"] <= 0 {
		fmt.Println("Si vous n'en avez plus, je vous offre la première, si je vous laisse partir sans potion et que vous périssez, cela pourrait nuire à la réputation de mon établissement")
	}
	fmt.Println("Potion de poison, 6 Pièces d'or: 2")
	fmt.Println("Livre de sort: Boule de feu, 25 Pièces d'or : 3")
	fmt.Println("Vendre: 4")
	fmt.Println("Page 2: 5")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case "1":
		if p.inventory["Potion de PV"] < 1 {
			p.AddDealer("Potion de PV", 1, 0)
			p.Dealer()
		} else {
			p.AddDealer("Potion de PV", 1, 3)
			p.Dealer()
		}
	case "2":
		p.AddDealer("Potion de poison", 1, 6)
		p.Dealer()
	case "3":
		p.AddDealer("Livre de sort: Boule de feu", 1, 25)
		p.Dealer()
	case "4":
		p.Dealer3()
	case "5":
		p.Dealer2()
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer()
	}
}
func (p *Player) Dealer2() {
	fmt.Println("----------------------------")
	fmt.Println("Page 2: Composants")
	fmt.Println("Votre or:", p.money)
	fmt.Println("------")
	fmt.Println("Fourrure de loup, 4 Pièces d'or: 1")
	fmt.Println("Peau de troll, 7 Pièces d'or: 2")
	fmt.Println("Cuir de sanglier, 3 Pièces d'or: 3")
	fmt.Println("Plume de corbeau, 1 Pièces d'or: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.AddDealer("Fourrure de loup", 1, 4)
		p.Dealer2()
	case "2":
		p.AddDealer("Peau de troll", 1, 7)
		p.Dealer2()
	case "3":
		p.AddDealer("Cuir de sanglier", 1, 3)
		p.Dealer2()
	case "4":
		p.AddDealer("Plume de corbeau", 1, 1)
		p.Dealer2()
	case "0":
		p.Dealer()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer2()
	}
}

/*func (p *Player) Dealer3() {
	fmt.Println("Page des ventes")
	fmt.Println("Vos objets:")
	for i:=0, i <  len(p.inventory) ; i++ {
		fmt.Println(nb, i)
		fmt.Println(i, ": 1")
		fmt.Println(i, ": 2")
		fmt.Println(i, ": 3")
		fmt.Println(i, ": 4")
		fmt.Println(i, ": 5")
		fmt.Println(i, ": 6")

		fmt.Println("------")
		fmt.Scanln(&input)
		switch input {
		case "1":
			p.SellInventory(i)
		case "2":
			p.SellInventory(i)
		case "3":
			p.SellInventory(i)
		case "4":
			p.SellInventory(i)
		case "5":
			p.SellInventory(i)
		case "6":
			p.SellInventory(i)
		}
	}
}
*/
