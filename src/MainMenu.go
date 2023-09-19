package ProjetRed

import "fmt"

func MainMenu(p *Player) {
	var input int
	fmt.Println("Menu principal:")
	fmt.Println("Information personnage: 1")
	fmt.Println("Inventaire: 2")
	fmt.Println("Quitter: 3")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.DisplayPlayerInfo()
	case 2: p.AccessInventory()
	case 3: fmt.Println("A plus !!")
	return
	default: fmt.Println("Erreur !!")
	}
}
