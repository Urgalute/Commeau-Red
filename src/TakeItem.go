package ProjetRed

import "fmt"

// Prendre une potion de PV
func (p *Player) TakePot() {
	if p.actuallife > p.maxlife-50 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion de PV"] > 0 {
		p.inventory["Potion de PV"]--
		p.actuallife = p.actuallife + 50
		fmt.Println("----------------------------")
		fmt.Println("Une potion de PV a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		p.DeleteInventory()
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de PV"])
		}
		fmt.Println("----------------------------")
	}
	p.AccessInventory()
}

// Potion de poison
func (p *Player) PoisonPot() {

}
