package ProjetRed

import (
	"fmt"
	"time"
)

// Prendre une potion de PV
func (p *Player) TakePot() {
	if p.actuallife > p.maxlife-50 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion de vie")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion de PV"] > 0 {
		p.inventory["Potion de PV"]--
		p.actuallife += 50
		fmt.Println("----------------------------")
		fmt.Println("Une potion de PV a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		p.DeleteInventory()
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de PV"])
		}
		fmt.Println("----------------------------")
	}
}

// Potion de poison
func (p *Player) PoisonPot() {
	if p.inventory["Potion de poison"] > 0 {
		fmt.Println("A l'avenir quand vous verrez un liquide vert, fûmant, dans une fiole en verre, évitez de le boire..")
		time.Sleep(3 * time.Second)
		for i := 0; i < 3; i++ {
			p.actuallife -= 10
			fmt.Println("Votre vie:", p.actuallife, "/", p.maxlife)
			time.Sleep(1 * time.Second)
		}
		p.inventory["Potion de poison"]--
		p.DeleteInventory()
		p.Wasted()
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
func (p *Player) TakeManaPot() {
	if p.actualmana > p.maxmana-25 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion de mana")
	}
	if p.actualmana <= p.maxmana-25 && p.inventory["Potion de mana"] > 0 {
		p.inventory["Potion de mana"]--
		p.actualmana += 25
		fmt.Println("----------------------------")
		fmt.Println("Une potion de mana a été utilisé, nouveau montant de mana:", p.actualmana, "/", p.maxmana)
		p.DeleteInventory()
		if p.inventory["Potion de mana"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de mana"])
		}
		fmt.Println("----------------------------")
	}
}
