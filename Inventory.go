package ProjetRed

import "fmt"

func (p player) AccessInventory() {
	fmt.Println("Votre inventaire:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
}

func (p *player) TakePot() {
	if p.inventory["Potion"] > 0 && p.actuallife <= p.maxlife-50 {
		p.inventory["Potion"]--
		p.actuallife = p.actuallife + 50
		fmt.Println("Une potion a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		fmt.Println("Il vous reste:", p.inventory["Potion"], "Potion")
	} else {
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
}
