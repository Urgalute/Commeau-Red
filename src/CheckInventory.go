package ProjetRed

import "fmt"

// Fonction qui vérifie la map p.inventory, afin de supprimer l'item et mettre un message si le personnage n'en a plus. A mettre après chaque décrémentation d'item

func (p *Player) ChechInventory() {
	for i := range p.inventory {
		if p.inventory[i] == 0 {
			fmt.Println("Vous n'avez plus de", i, "!")
			delete(p.inventory, i)
		}
	}
	fmt.Println("----------------------------")
}
