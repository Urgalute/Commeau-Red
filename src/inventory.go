package ProjetRed

import (
	"fmt"
	"time"
)

// Inventaire de combat avec les actions associées
func (p *Player) FightInventory() {
	fmt.Println("Menu inventaire:")
	fmt.Println("------")
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	if p.inventory["Potion de PV"] > 0 {
		fmt.Println("1: Utiliser une potion de PV (+50 PV)")
	}
	if p.inventory["Potion de mana"] > 0 {
		fmt.Println("2: Utiliser une potion de mana (+25 PM)")
		fmt.Println("0: Retour")
		fmt.Println("----------------------------")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.inventory["Potion de PV"] > 0 {
				if p.TakePot(); true {
					p.AttackGoblin(Turn)
				} else {
					p.FightInventory()
				}
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.FightInventory()
			}
		case "2":
			if p.inventory["Potion de mana"] > 0 {
				if p.TakeManaPot(); true {
					p.AttackGoblin(Turn)
				} else {
					p.FightInventory()
				}
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.FightInventory()
			}
		case "0":
			p.TrainingFight(Turn)
		default:
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.FightInventory()
		}
	}
}
func (p *Player) AccessInventory() {
	//Affiche l'inventaire et les actions possible
	var input string
	fmt.Println("Menu inventaire:")
	fmt.Println("------")
	fmt.Println("Votre argent:", p.money["Pièces d'or"], "Pièces d'or")
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	if p.inventory["Potion de PV"] > 0 {
		fmt.Println("1: Utiliser une potion de PV (+50 PV)")
	}
	if p.inventory["Potion de mana"] > 0 {
		fmt.Println("2: Utiliser une potion de mana (+25 PM)")
	}
	if p.inventory["Potion de poison"] > 0 {
		fmt.Println("3: Utiliser une potion de poison (Inflige 10 dégâts par seconde pendant 3 secondes)")
	}
	if p.inventory["Chapeau de l'aventurier"] > 0 && p.equipment.head != "Chapeau de l'aventurier" {
		fmt.Println("4: Equiper le chapeau de l'aventurier (+10 PV et PM max)")
	}
	if p.inventory["Tunique de l'aventurier"] > 0 && p.equipment.torse != "Tunique de l'aventurier" {
		fmt.Println("5: Equiper la tunique de l'aventurier (+25 PV et PM max)")
	}
	if p.inventory["Bottes de l'aventurier"] > 0 && p.equipment.foot != "Bottes de l'aventurier" {
		fmt.Println("6: Equiper les bottes de l'aventurier (+15 PV et PM max)")
	}
	if p.inventory["Livre de sort: Boule de feu"] > 0 {
		fmt.Println("7: Apprendre le sort Boule de feu")
	}
	fmt.Println("0: Retour")
	fmt.Println("----------------------------")
	fmt.Scanln(&input)
	switch input {
	case "1":
		if p.inventory["Potion de PV"] > 0 {
			p.TakePot()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "3":
		if p.inventory["Potion de poison"] > 0 {
			fmt.Println("A l'avenir quand vous verrez un liquide vert, fûmant, dans une fiole en verre, évitez de le boire..")
			time.Sleep(3 * time.Second)
			p.PoisonPot()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "2":
		if p.inventory["Potion de mana"] > 0 {
			p.TakeManaPot()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "4":
		if p.equipment.head == "Chapeau de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Chapeau de l'aventurier"] > 0 {
			p.equipment.head = "Chapeau de l'aventurier"
			p.maxlife += 10
			p.maxmana += 10
			p.inventory["Chapeau de l'aventurier"]--
			fmt.Println("Vous avez bien équipé le chapeau de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "5":
		if p.equipment.torse == "Tunique de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Tunique de l'aventurier"] > 0 {
			p.equipment.torse = "Tunique de l'aventurier"
			p.maxlife += 25
			p.maxmana += 25
			p.inventory["Tunique de l'aventurier"]--
			fmt.Println("Vous avez bien équipé la tunique de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "6":
		if p.equipment.foot == "Bottes de l'aventurier" {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Bottes de l'aventurier"] > 0 {
			p.equipment.foot = "Bottes de l'aventurier"
			p.maxlife += 15
			p.maxmana += 15
			p.inventory["Bottes de l'aventurier"]--
			fmt.Println("Vous avez bien équipé les bottes de l'aventurier")
			p.DeleteInventory()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "7":
		if p.inventory["Livre de sort: Boule de feu"] > 0 {
			p.SpellBook()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
		p.AccessInventory()
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.AccessInventory()
	}
}
func (p *Player) AddDealer(i string, j int, k int) {
	//Add les achats au marchand dans la map p.inventory
	if p.money["Pièces d'or"] >= k {
		if len(p.inventory) >= p.rangeinventory {
			fmt.Println("Votre inventaire est plein !", "\n", "Pensez à vous équiper de vos objet, ou vendez votre surplus au marchand !")
		} else {
			p.inventory[i] += j
			p.money["Pièces d'or"] -= k
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
	} else {
		p.inventory[i] += j
		p.money["Pièces d'or"] -= k
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
	//Vérifie la place disponible dans l'inventaire (ne compte que les items unique, pas leurs quantité pour le moment)
	for i := range p.inventory {
		if p.inventory[i] >= p.rangeinventory {
			fmt.Println("Vous n'avez plus de place dans votre inventaire")
		}
	}
}
func (p *Player) SellInventory(i string) {
	//Après la vente, definie un pris de vente et l'ajoute à mon or, + message l'indiquant
	var j int
	p.inventory[i]--
	switch i {
	case "Potion de PV":
		j = 1
	case "Potion de poison":
		j = 7
	case "Potion de mana":
		j = 3
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
	p.money["Pièces d'or"] += j
}
