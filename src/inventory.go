package ProjetRed

import "fmt"

func (p *Player) AccessInventory() {
	//Affiche l'inventaire et les actions possible
	var input string
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
		case "1":
			p.TakePot()
		case "0":
			p.MainMenu()
		default:
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.AccessInventory()
		}
	} else {
		fmt.Println("Retour: 0")
		fmt.Scanln(&input)
		switch input {
		case "0":
			p.MainMenu()
		default:
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.AccessInventory()
		}
	}
}
func (p *Player) AddDealer(i string, j int, k int) {
	//Add les achats au marchand dans la map p.inventory
	if p.money >= k {
		if len(p.inventory) >= 10 {
			fmt.Println("Votre inventaire est plein !", "\n", "Pensez à vous équiper de vos objet, ou vendez votre surplus au marchand !")
		} else {
			p.inventory[i] += j
			p.money -= k
			fmt.Println("----------------------------")
			fmt.Println("Vous avez acheté:", j, i)
			fmt.Println("Total dans votre inventaire:", p.inventory[i])
			fmt.Println("----------------------------")
		}
	} else {
		fmt.Println("Désolé mais vous n'avez pas assez d'or, vous avez seulement", p.money, "pièces d'or sur les", k, "demandé pour cet objet")
	}
}
func (p *Player) AddCraft(i string, j int, k int) {
	//Add les craft du forgeron dans la map p.inventory
	p.inventory[i] += j
	p.money -= k
	fmt.Println("----------------------------")
	fmt.Println("Vous venez de fabriquer: 1 ", i)
	fmt.Println("Total dans votre inventaire:", p.inventory[i])
	fmt.Println("----------------------------")
}
func (p *Player) DeleteInventory() {
	//Vérifie la map p.inventory, supp l'item si := 0 et mettre un message si le personnage n'en a plus. A mettre après chaque décrémentation d'item
	for i := range p.inventory {
		if p.inventory[i] == 0 {
			fmt.Println("Vous n'avez plus de", i, "!")
			delete(p.inventory, i)
		}
	}
	fmt.Println("----------------------------")
}
func (p *Player) CheckInventory() {
	for i := range p.inventory {
		if p.inventory[i] >= 10 {
			fmt.Println("Vous n'avez plus de place dans votre inventaire")
			return
		}
	}
}
func (p *Player) SellInventory(i string) {
	var j int
	p.inventory[i]-- 
	 switch i {
	 case "Potion de PV": j = 1
	 case "Potion de poison": j = 2
	 case "Livre de sort: Boule de feu": j = 10
	 case "Fourrure de loup": j = 2
	 case "Peau de troll": j = 3
	 case "Cuir de sanglier": j = 1
	 case "Plume de corbeau": j = 0
	 case "Chapeau de l'aventurier": j = 12
	 case "Tunique de l'aventurier": j = 30
	 case "Bottes de l'aventurier": j = 18
	 }
	if p.inventory[i] == 0 {
		fmt.Println("Vous n'avez plus de", i, "!")
		delete(p.inventory, i)
	}
	p.money += j
}
