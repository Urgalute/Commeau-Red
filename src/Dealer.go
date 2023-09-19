package ProjetRed

import (
	"fmt"
)

func (p *Player) Dealer() {
	var input int
	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("Acheter une potion de vie: 1")
	fmt.Println("Retour au menu principal: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.AddPot()
	case 0:
		p.MainMenu()
	}
}
