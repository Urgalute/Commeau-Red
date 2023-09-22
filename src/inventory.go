package ProjetRed

import "fmt"

func (p *Player) AccessInventory() {
	//Affiche l'inventaire et les actions possible
	var input string
	fmt.Println("Menu inventaire:")
	fmt.Println("------")
	fmt.Println("Votre argent:", p.money, "Pièces d'or")
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	if p.inventory["Potion de PV"] > 0 {
		fmt.Println("Utiliser une potion de PV: 1")
	}
	if p.inventory["Chapeau de l'aventurier"] > 0 && p.equipment.head != "Chapeau de l'aventurier" {
		fmt.Println("Equiper le chapeau de l'aventurier: 2")
	}
	if p.inventory["Tunique de l'aventurier"] > 0 && p.equipment.torse != "Tunique de l'aventurier" {
		fmt.Println("Equiper la tunique de l'aventurier: 3")
	}
	if p.inventory["Bottes de l'aventurier"] > 0 && p.equipment.foot != "Bottes de l'aventurier" {
		fmt.Println("Equiper les bottes de l'aventurier: 4")
	}
	fmt.Println("Retour: 0")
	fmt.Println("----------------------------")
	fmt.Scanln(&input)
	switch input {
	case "1":
		if p.inventory["Potion de PV"] > 0 {
			p.TakePot()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "2":
		if p.equipment.head == "Chapeau de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Chapeau de l'aventurier"] > 0 {
			p.equipment.head = "Chapeau de l'aventurier"
			p.maxlife += 10
			p.inventory["Chapeau de l'aventurier"]--
			fmt.Println("Vous avez bien équipé le chapeau de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "3":
		if p.equipment.torse == "Tunique de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Tunique de l'aventurier"] > 0 {
			p.equipment.torse = "Tunique de l'aventurier"
			p.maxlife += 25
			p.inventory["Tunique de l'aventurier"]--
			fmt.Println("Vous avez bien équipé la tunique de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "4":
		if p.equipment.foot == "Bottes de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Bottes de l'aventurier"] > 0 {
			p.equipment.foot = "Bottes de l'aventurier"
			p.maxlife += 15
			p.inventory["Bottes de l'aventurier"]--
			fmt.Println("Vous avez bien équipé les bottes de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.AccessInventory()
	}
}
func (p *Player) AddDealer(i string, j int, k int) {
	//Add les achats au marchand dans la map p.inventory
	if p.money >= k {
		if len(p.inventory) >= p.rangeinventory {
			fmt.Println("Votre inventaire est plein !", "\n", "Pensez à vous équiper de vos objet, ou vendez votre surplus au marchand !")
			p.Dealer()
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
	if len(p.inventory) >= p.rangeinventory {
		fmt.Println("Votre inventaire est plein !", "\n", "Pensez à vous équiper de vos objet, ou vendez votre surplus au marchand !")
		p.BlackSmith()
	} else {
		p.inventory[i] += j
		p.money -= k
		fmt.Println("----------------------------")
		fmt.Println("Vous venez de fabriquer: 1 ", i)
		fmt.Println("Total dans votre inventaire:", p.inventory[i])
		fmt.Println("----------------------------")
	}
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
		if p.inventory[i] >= p.rangeinventory {
			fmt.Println("Vous n'avez plus de place dans votre inventaire")
		}
	}
}
func (p *Player) SellInventory(i string) {
	var j int
	p.inventory[i]--
	switch i {
	case "Potion de PV":
		j = 1
	case "Potion de poison":
		j = 2
	case "Livre de sort: Boule de feu":
		j = 10
	case "Fourrure de loup":
		j = 2
	case "Peau de troll":
		j = 3
	case "Cuir de sanglier":
		j = 1
	case "Plume de corbeau":
		j = 0
	case "Chapeau de l'aventurier":
		j = 7
	case "Tunique de l'aventurier":
		j = 17
	case "Bottes de l'aventurier":
		j = 10
	default:
		fmt.Println("Erreur")
	}
	fmt.Println("Vous avez vendu 1", i, "pour", j, "Pièces d'or")
	p.DeleteInventory()
	p.money += j
}
