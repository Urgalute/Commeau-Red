package ProjetRed

import "fmt"

func (p *Player) AddInventory(i string, j int, k int) {
	if p.money >= k {
		p.inventory[i] += j
		p.money -= k
		fmt.Println("----------------------------")
		fmt.Println("Vous avez acheté:", j, i)
		fmt.Println("Total dans votre inventaire:", p.inventory[i])
		fmt.Println("----------------------------")
	} else {
		fmt.Println("Désolé mais vous n'avez pas assez d'or, vous avez seulement", p.money, "pièces d'or sur les", k, "demandé pour cet objet")
	}
}
func (p *Player) AddCraft(i string, j int, k int) {
	p.inventory[i] += j
	p.money -= k
	fmt.Println("----------------------------")
	fmt.Println("Vous venez de fabriquer: 1 ", i)
	fmt.Println("Total dans votre inventaire:", p.inventory[i])
	fmt.Println("----------------------------")
}
