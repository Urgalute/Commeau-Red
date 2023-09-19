package ProjetRed

import "fmt"

func (p *Player) AddPot() {
	p.AddInventory("Potion", 1)
	p.money -= 3
	fmt.Println("----------------------------")
	p.Dealer()
}
func (p *Player) AddWolfFur() {
	p.AddInventory("Fourrure de loup", 1)
	p.money -= 4
	fmt.Println("----------------------------")
	p.Dealer()
}
func (p *Player) TrollSkin() {
	p.AddInventory("Peau de troll", 1)
	p.money -= 7
	fmt.Println("----------------------------")
	p.Dealer()
}
func (p *Player) BoarLeather() {
	p.AddInventory("Cuir de sanglier", 1)
	p.money -= 3
	fmt.Println("----------------------------")
	p.Dealer()
}
func (p *Player) RavenFeather() {
	p.AddInventory("Plume de corbeau", 1)
	p.money -= 1
	fmt.Println("----------------------------")
	p.Dealer()
}
