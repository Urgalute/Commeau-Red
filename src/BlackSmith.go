package ProjetRed

import "fmt"

func (p *Player) BlackSmith() {
	fmt.Println("Chapeau de l'aventurier: 1")
	fmt.Println("Tunique de l'aventurier: 2")
	fmt.Println("Bottes de l'aventurier: 3")
	fmt.Scanln(&input)
	switch input {
	case 1: 
	if p.inventory["Plume de corbeau"] > 0 && p.inventory["Cuir de sanglier"] > 0 {
		p.AddInventory("Chapeau de l'aventurier", 1, 0)
		p.inventory["Plume de corbeau"]--
		p.inventory["Cuir de sanglier"]--
		fmt.Println("Vous venez de fabriquer: 1 Chapeau de l'aventurier")
	}
}
}