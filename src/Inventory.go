package ProjetRed

import "fmt"

func (p *Player) AccessInventory() {
	fmt.Println("Votre inventaire:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
		if p.inventory[i] == 0 {
			delete(p.inventory, i)
		}
	}
}

func (p *Player) TakePot() {
	if p.actuallife > p.maxlife-50 {
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion"] > 0 {
		p.inventory["Potion"]--
		p.actuallife = p.actuallife + 50
		fmt.Println("Une potion a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		if p.inventory["Potion"] <= 0 {
			delete(p.inventory, "Potion")
			fmt.Println("Vous n'avez plus de potion !")
		} else {
			fmt.Println("Il vous reste:", p.inventory["Potion"], "Potion")
		}
	}
}
