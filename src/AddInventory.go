package ProjetRed

import "fmt"

func (p *Player) AddInventory(i string, j int) {
	p.inventory[i] += j
	fmt.Println("Vous avez acheté:", j, i)
	fmt.Println("Total dans votre inventaire:", p.inventory[i], i)
}
