package ProjetRed

import (
	"fmt"
)

func (p *Player) Dealer() {
	var input int
	fmt.Println("----------------------------")
	fmt.Println("Acheter une potion de vie: 1")
	fmt.Println("Fourrure de loup: 2")
	fmt.Println("Peau de troll: 3")
	fmt.Println("Cuir de sanglier: 4")
	fmt.Println("Plume de corbeau: 5")
	fmt.Println("Retour au menu principal: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddPot()
	case 2:
		p.AddWolfFur()
	case 3:
		p.TrollSkin()
	case 4:
		p.BoarLeather()
	case 5:
		p.RavenFeather()
	case 0:
		p.MainMenu()
	}
}
