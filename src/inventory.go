package ProjetRed

import (
	"fmt"
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
	}
	if p.inventory["Potion de poison"] > 0 {
		fmt.Println("3: Utiliser une potion de poison (Inflige 10 dégât par seconde pendant 3 secondes)")
	}
	if p.inventory["Marteau de forgeron"] > 0 {
		fmt.Println("4: Utiliser un marteau de forgeron (Répare votre équipement)")
	}
	fmt.Println("0: Retour")
	fmt.Println("----------------------------")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		if p.inventory["Potion de PV"] > 0 {
			if p.actuallife <= p.maxlife-50 {
				p.TakePot()
				p.AttackGoblin(Turn)
			} else {
				fmt.Println("Vous ne pouvez pas prendre de potion de vie")
				p.FightInventory()
			}
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.FightInventory()
		}
	case "2":
		ClearTerminal()
		if p.inventory["Potion de mana"] > 0 {
			if p.actualmana <= p.maxmana-25 {
				p.TakeManaPot()
				p.AttackGoblin(Turn)
			} else {
				fmt.Println("Vous ne pouvez pas prendre de potion de vie")
				p.FightInventory()
			}
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.FightInventory()
		}
	case "3":
		ClearTerminal()
		if p.inventory["Potion de poison"] > 0 {
			if p.actuallife <= p.actuallife-30 {
				fmt.Println("Vous êtes sûr de vous ?")
				fmt.Println("1: Oui")
				fmt.Println("0: Non")
				fmt.Scanln(&input)
				switch input {
				case "1":
					ClearTerminal()
					p.PoisonPot()
					p.AttackGoblin(Turn)
				case "0":
					ClearTerminal()
					p.FightInventory()
				default:
					ClearTerminal()
					fmt.Println("---------------------------------------------------------------------------------------------------------")
					fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
					fmt.Println("---------------------------------------------------------------------------------------------------------")
					p.FightInventory()
				}
			} else {
				ClearTerminal()
				p.PoisonPot()
				p.AttackGoblin(Turn)
			}
		} else {
			ClearTerminal()
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.FightInventory()
		}
	case "4":		
	ClearTerminal()
				p.TakeHammer()
				p.FightInventory()
	case "0":
		ClearTerminal()
		p.TrainingFight(Turn)
	default:
		ClearTerminal()
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.FightInventory()
	}
}
func (p *Player) AccessInventory() {
	//Affiche l'inventaire et les actions possible
	var input string
	fmt.Println("Menu inventaire:")
	fmt.Println("------")
	fmt.Println("Votre argent:")
	for i := range p.money {
		fmt.Println(p.money[i], i)
	}
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
	if p.inventory["Marteau de forgeron"] > 0 {
		fmt.Println("4: Utiliser un marteau de forgeron (Répare votre équipement)")
	}
	if p.inventory["Chapeau de l'aventurier"] > 0 && p.equipment.head["Chapeau de l'aventurier"] <= 0 {
		fmt.Println("5: Equiper le chapeau de l'aventurier (+10 PV et PM max)")
	}
	if p.inventory["Tunique de l'aventurier"] > 0 && p.equipment.torse["Tunique de l'aventurier"] <= 0 {
		fmt.Println("6: Equiper la tunique de l'aventurier (+25 PV et PM max)")
	}
	if p.inventory["Bottes de l'aventurier"] > 0 && p.equipment.foot["Bottes de l'aventurier"] <= 0 {
		fmt.Println("7: Equiper les bottes de l'aventurier (+15 PV et PM max)")
	}
	if p.inventory["Livre de sort: Boule de feu"] > 0 {
		fmt.Println("8: Apprendre le sort Boule de feu")
	}
	fmt.Println("0: Retour")
	fmt.Println("----------------------------")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		if p.inventory["Potion de PV"] > 0 {
			p.TakePot()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "3":
		ClearTerminal()
		if p.inventory["Potion de poison"] > 0 {
			p.PoisonPot()
			if p.Wasted() {
				p.AccessInventory()
			}
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "2":
		ClearTerminal()
		if p.inventory["Potion de mana"] > 0 {
			p.TakeManaPot()
			p.AccessInventory()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "4":
		ClearTerminal()
		p.TakeHammer()
	p.AccessInventory()
	case "5":
		ClearTerminal()
		if p.equipment.head["Chapeau de l'aventurier"] == -5 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.equipment.head["Chapeau de l'aventurier"] > 0 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Chapeau de l'aventurier"] > 0 {
			p.equipment.head["Chapeau de l'aventurier"] = 50
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
	case "6":
		ClearTerminal()
		if p.equipment.torse["Tunique de l'aventurier"] == -5 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.equipment.torse["Tunique de l'aventurier"] > 0 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Tunique de l'aventurier"] > 0 {
			p.equipment.torse["Tunique de l'aventurier"] = 50
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
	case "7":
		ClearTerminal()
		if p.equipment.foot["Bottes de l'aventurier"] == -5 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.equipment.foot["Bottes de l'aventurier"] > 0 {
			fmt.Println("Un autre exemplaire de cet objet est déjà équipé")
			p.AccessInventory()
		} else if p.inventory["Bottes de l'aventurier"] > 0 {
			p.equipment.foot["Bottes de l'aventurier"] = 50
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
	case "8":
		ClearTerminal()
		if p.inventory["Livre de sort: Boule de feu"] > 0 {
			p.SpellBook()
		} else {
			fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
			p.AccessInventory()
		}
	case "0":
		ClearTerminal()
		p.MainMenu()
	default:
		ClearTerminal()
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
	//Après la vente, definie un pris de vente et l'ajoute à mon or, + message
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
func (p *Player) CheckEquipment(rand int, dammage int) {
	if p.equipment.head["Chapeau de l'aventurier"] > 0 {
		if rand > 0 && rand <= 33 {
			p.equipment.head["Chapeau de l'aventurier"] -= dammage
			for i := range p.equipment.head {
				fmt.Println("Vous perdez", dammage, "de durabilité sur votre Chapeau de l'aventurier", p.equipment.head[i], "/ 50)")
			}
			if p.equipment.head["Chapeau de l'aventurier"] == 0 {
				p.equipment.head["Chapeau de l'aventurier"] -= 5
			}
			if p.equipment.head["Chapeau de l'aventurier"] < 0 {
				fmt.Println("La durabilité de Chapeau de l'aventurier est à 0, vous perdez le bonus de statistique associé")
				p.maxlife -= 10
				p.maxmana -= 10
			}
		}
	}
	if p.equipment.torse["Tunique de l'aventurier"] > 0 {
		if rand > 33 && rand <= 66 {
			p.equipment.torse["Tunique de l'aventurier"] -= dammage
			for i := range p.equipment.torse {
				fmt.Println("Vous perdez", dammage, "de durabilité sur votre Tunique de l'aventurier", p.equipment.torse[i], "/ 50)")
			}
			if p.equipment.torse["Tunique de l'aventurier"] == 0 {
				p.equipment.torse["Tunique de l'aventurier"] -= 5
			}
			if p.equipment.torse["Tunique de l'aventurier"] < 0 {
				fmt.Println("La durabilité de votre Tunique de l'aventurier est à 0, vous perdez le bonus de statistique associé")
				p.maxlife -= 25
				p.maxmana -= 25
			}

		}
	}
	if p.equipment.foot["Bottes de l'aventurier"] > 0 {
		if rand > 66 && rand <= 99 {
			p.equipment.foot["Bottes de l'aventurier"] -= dammage
			for i := range p.equipment.foot {
				fmt.Println("Vous perdez", dammage, "de durabilité sur vos Bottes de l'aventurier", p.equipment.foot[i], "/ 50)")
			}
			if p.equipment.foot["Bottes de l'aventurier"] == 0 {
				p.equipment.foot["Bottes de l'aventurier"] -= 5
			}
			if p.equipment.foot["Bottes de l'aventurier"] < 0 {
				fmt.Println("La durabilité de vos Bottes de l'aventurier est à 0, vous perdez le bonus de statistique associé")
				p.maxlife -= 15
				p.maxmana -= 15
			}
		}
	}
}
