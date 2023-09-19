package ProjetRed

import "fmt"

func (p *Player) MainMenu() {
	var input int
	fmt.Println("----------------------------")
	fmt.Println("Menu principal:")
	fmt.Println("------")
	fmt.Println("Information personnage: 1")
	fmt.Println("Inventaire: 2")
	fmt.Println("Marchand: 3")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.DisplayPlayerInfo()
	case 2:
		p.AccessInventory()
	case 3:
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le marchand !")
		p.Dealer()
	case 0:
		fmt.Println("A plus !")
	}
}
