package ProjetRed

import "fmt"

func (p *Player) AddPot() {
	p.inventory["Potion"]++
	fmt.Println("Vous avez acheter 1 Potion")

}
