package ProjetRed

import (
	"fmt"
	"time"
)

// Prendre une potion de PV
func (p *Player) TakePot() bool {
	i := false
	if p.actuallife > p.maxlife-50 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion de PV"] > 0 {
		p.inventory["Potion de PV"]--
		p.actuallife += 50
		fmt.Println("----------------------------")
		fmt.Println("Une potion de PV a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		p.DeleteInventory()
		i = true
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de PV"])
		}
		fmt.Println("----------------------------")
	}
	return i
}

// Potion de poison
func (p *Player) PoisonPot() {
	if p.inventory["Potion de poison"] > 0 {
		for i := 0; i < 3; i++ {
			p.actuallife -= 10
			fmt.Println("Votre vie:", p.actuallife, "/", p.maxlife)
			time.Sleep(1 * time.Second)
		}
		p.inventory["Potion de poison"]--
		p.DeleteInventory()
	}
}

// Livre de sort: Boule de feu
func (p *Player) SpellBook() {
	if p.inventory["Livre de sort: Boule de feu"] > 0 {
		p.spell[1] = "Boule de feu"
		p.inventory["Livre de sort: Boule de feu"]--
		p.DeleteInventory()
		fmt.Println("Vous avez appris un nouveau sort: Boule de feu")
	}
}

// Potion de mana
func (p *Player) TakeManaPot() bool {
	i := false
	if p.actualmana > p.maxmana-25 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion")
	}
	if p.actualmana <= p.maxmana-50 && p.inventory["Potion de mana"] > 0 {
		p.inventory["Potion de mana"]--
		p.actualmana += 25
		fmt.Println("----------------------------")
		fmt.Println("Une potion de mana a été utilisé, nouveau montant de mana:", p.actualmana, "/", p.maxmana)
		p.DeleteInventory()
		i = true
		if p.inventory["Potion de mana"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de mana"])
		}
		fmt.Println("----------------------------")
	}
	return i
}
