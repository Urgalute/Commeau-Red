package Red

import (
	"fmt"
)

func (p player) AccessInventory() {

	for i := range p.inventory {
		fmt.Println(i, "=", p.inventory[i])
	}
}

func (p player) TakePot() {
	if p.inventory["Potion"] > 0 {
		p.inventory["Potion"]--

		if p.pvactual < p.pvmax -50 {
			p.pvactual = +50
		} else {
			fmt.Println("Vous ne pouvez pas")
			
		}
	}
}
