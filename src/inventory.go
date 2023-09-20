package ProjetRed

import "fmt"

func (p *Player) AccessInventory() {

	var input int
	fmt.Println("Menu inventaire:")
	fmt.Println("------")
	fmt.Println("Votre argent:", p.money, "Pièces d'or")
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	fmt.Println("----------------------------")
	if p.inventory["Potion de PV"] > 0 {
		fmt.Println("Utiliser une potion de PV: 1")
		fmt.Println("Retour: 0")
		fmt.Scanln(&input)
		switch input {
		case 1:
			p.TakePot()
		case 0:
			p.MainMenu()
		}
	}
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case 2:
	case 0:
		p.MainMenu()
	}

}

func (p *Player) TakePot() {
	if p.actuallife > p.maxlife-50 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion de PV"] > 0 {
		p.inventory["Potion de PV"]--

		p.actuallife = p.actuallife + 50
		fmt.Println("----------------------------")
		p.ChechInventory()
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Une potion de PV a été utilisé, il vous en reste:", p.inventory["Potion de PV"])
		}
		fmt.Println("Nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		fmt.Println("----------------------------")
	}

	p.AccessInventory()
}
