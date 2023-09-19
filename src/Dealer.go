package ProjetRed

import (
	"fmt"
)

var input int

func (p *Player) Dealer() {
	fmt.Println("----------------------------")
	fmt.Println("Acheter une potion de vie: 1")
	fmt.Println("Potion de poison: 2")
	fmt.Println("Livre de sort, Boule de feu: 3")
	fmt.Println("Page 2: 4")
	fmt.Println("Retour au menu principal: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddInventory("Potion de vie", 1, 3)
	case 2:
		p.AddInventory("Potion de poison", 1, 6)
	case 3:
		p.AddInventory("Livre de sort: Boule de feu", 1, 25)
	case 4:
		p.Dealer2()
	case 0:
		p.MainMenu()
	}
}
func (p *Player) Dealer2() {
	fmt.Println("----------------------------")
	fmt.Println("Fourrure de loup: 1")
	fmt.Println("Peau de troll: 2")
	fmt.Println("Cuir de sanglier: 3")
	fmt.Println("Plume de corbeau: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddInventory("Peau de troll", 1, 7)
	case 2:
		p.AddInventory("Cuir de sanglier", 1, 3)
	case 3:
		p.AddInventory("Plume de corbeau", 1, 1)
	case 0:
		p.Dealer()
	}
}
