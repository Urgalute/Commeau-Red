package ProjetRed

import "fmt"

func (p *Player) MainMenu() {
	var input int
	fmt.Println("Menu principal:")
	fmt.Println("Information personnage: 1")
	fmt.Println("Inventaire: 2")
	fmt.Println("Marchand: 3")
	fmt.Println("Quitter: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.DisplayPlayerInfo()
	case 2:
		p.AccessInventory()
	case 3:
		p.Dealer()
	case 0:
		return
	}
}
