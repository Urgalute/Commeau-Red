package ProjetRed

import "fmt"

func (p *Player) AccessInventory() {
	var input int
	fmt.Println("Votre inventaire:")
	for i := range p.inventory {
		if p.inventory[i] == 0 {
			delete(p.inventory, i)
		}
		fmt.Println(p.inventory[i], i)

	}
	fmt.Println("Utiliser une potion: 1")
	fmt.Println("Retour au menu principal: 0")
	fmt.Scanln(&input)
	switch input {
	case 1:
		p.TakePot()
	case 2:
	case 0:
		p.MainMenu()
		return
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
		p.AccessInventory()
			if p.inventory["Potion"] == 0 {
		delete(p.inventory, "Potion")
		fmt.Println("Vous n'avez plus de potion !")
	} else {
		fmt.Println("Il vous reste:", p.inventory["Potion"], "Potion")
	}
	p.AccessInventory()
}
}