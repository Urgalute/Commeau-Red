package ProjetRed

import (
	"fmt"
)

var input string

func (p *Player) Dealer() {
	//Marchand page 1 Utilitaires
	fmt.Println("Page 1: Utilitaires")
	fmt.Println("------")
	fmt.Println("Votre or:")
	for i := range p.money {
		fmt.Println(p.money[i], i)
	}
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("1: Acheter une potion de PV (+50 PV), 3 Pièces d'or")
	if p.inventory["Potion de PV"] <= 0 {
		fmt.Println("Si vous n'en avez plus, je vous offre la première, si je vous laisse partir sans potion et que vous périssez, cela pourrait nuire à la réputation de mon établissement")
	}
	fmt.Println("2: Acheter une potion de poison, 10 Pièces d'or")
	fmt.Println("3: Acheter une potion de mana (+25 PM), 5 Pièces d'or")
	fmt.Println("4: Acheter un marteau de forgeron (Vous permet de réparer votre équipement vous même), 20 Pièces d'or")
	if p.spell[1] != "Boule de feu" && p.inventory["Livre de sort: Boule de feu"] <= 0 {
		fmt.Println("5: Acheter le Livre de sort: Boule de feu, 25 Pièces d'or")
	}
	if p.rangeinventory == 10 {
		fmt.Println("6: Capacité d'inventaire +10 (3 Achats restant): 30 Pièces d'or")
	}
	if p.rangeinventory == 20 {
		fmt.Println("6: Capacité d'inventaire +10 (2 Achats restant): 30 Pièces d'or")
	}
	if p.rangeinventory == 30 {
		fmt.Println("6: Capacité d'inventaire +10 (1 Achats restant): 30 Pièces d'or")
	}
	if p.rangeinventory == 40 {
		fmt.Println("Capacité d'inventaire +10 (Epuisé, limite d'inventaire max atteinte)")
	}
	fmt.Println("7: Réparer l'équipement (10 Pièces d'or)")
	fmt.Println("8: Vendre")
	fmt.Println("9: Page 2, composants")
	fmt.Println("0: Retour")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		if p.inventory["Potion de PV"] < 1 {
			p.AddDealer("Potion de PV", 1, 0)
			p.Dealer()
		} else if p.inventory["Potion de PV"] < p.rangeinventory {
			p.AddDealer("Potion de PV", 1, 3)
			p.Dealer()
		} else {
			fmt.Println("Limite d'inventaire atteinte")
			p.Dealer()
		}
	case "2":
		ClearTerminal()
		if p.inventory["Potion de poison"] < p.rangeinventory {
			p.AddDealer("Potion de poison", 1, 10)
			p.Dealer()
		} else {
			fmt.Println("Limite d'inventaire atteinte")
			p.Dealer()
		}
	case "3":
		ClearTerminal()
		if p.inventory["Potion de mana"] < p.rangeinventory {
			p.AddDealer("Potion de mana", 1, 5)
			p.Dealer()
		} else {
			fmt.Println("Limite d'inventaire atteinte")
			p.Dealer()
		}
	case "4":
		ClearTerminal()
		if p.inventory["Marteau de forgeron"] < p.rangeinventory {
			p.AddDealer("Marteau de forgeron", 1, 20)
			p.Dealer()
		} else {
			fmt.Println("Limite d'inventaire atteinte")
			p.Dealer()
		}
	case "5":
		ClearTerminal()
		if p.inventory["Livre de sort: Boule de feu"] > 0 {
			fmt.Println("Vous avez déjà acheté ce sort")
			p.Dealer()
		} else if p.spell[1] == "Boule de feu" {
			fmt.Println("Vous connaissez déjà ce sort")
			p.Dealer()
		} else {
			p.AddDealer("Livre de sort: Boule de feu", 1, 25)
			p.Dealer()
		}
	case "6":
		ClearTerminal()
		if p.rangeinventory <= 30 {
			p.rangeinventory += 10
			fmt.Println("La taille maximale de votre inventaire passe à", p.rangeinventory, "emplacement !")
			p.Dealer()
		} else {
			fmt.Println("Bien essayé mais vous avez déja atteint la capacité d'inventaire max")
			p.Dealer()
		}
	case "7":
			ClearTerminal()
			p.TakeHammer2()

			p.Dealer()
	case "8":
		ClearTerminal()
		p.SellDealer()
	case "9":
		ClearTerminal()
		p.Dealer2()
	case "0":
		ClearTerminal()
		p.MainMenu()
	default:
		ClearTerminal()
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer()
	}
}
func (p *Player) Dealer2() {
	// Marchand page 2 Composants
	fmt.Println("----------------------------")
	fmt.Println("Page 2: Composants")
	fmt.Println("Votre or:")
	for i := range p.money {
		fmt.Println(p.money[i], i)}
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("------")
	fmt.Println("1: Fourrure de loup, 4 Pièces d'or")
	fmt.Println("2: Peau de troll, 7 Pièces d'or")
	fmt.Println("3: Cuir de sanglier, 3 Pièces d'or")
	fmt.Println("4: Plume de corbeau, 1 Pièces d'or")
	fmt.Println("0: Retour")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		p.AddDealer("Fourrure de loup", 1, 4)
		p.Dealer2()
	case "2":
		ClearTerminal()
		p.AddDealer("Peau de troll", 1, 7)
		p.Dealer2()
	case "3":
		ClearTerminal()
		p.AddDealer("Cuir de sanglier", 1, 3)
		p.Dealer2()
	case "4":
		ClearTerminal()
		p.AddDealer("Plume de corbeau", 1, 1)
		p.Dealer2()
	case "0":
		ClearTerminal()
		p.Dealer()
	default:
		ClearTerminal()
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer2()
	}
}

func (p *Player) SellDealer() {
	//Marchand menu vente, seulement si des items sont disponible
	var input string
	fmt.Println("Page des ventes:")
	fmt.Println("Votre or:")
	for i := range p.money {
		fmt.Println(p.money[i], i)
	}
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	fmt.Println("------")
	if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
		fmt.Println("Page 1, utilitaires:")
		fmt.Println("Votre or:")
		for i := range p.money {
			fmt.Println(p.money[i], i)
		}
		fmt.Println("------")
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("1: Vendre 1 Potion de PV (1 Pièce d'or)")
		}
		if p.inventory["Potion de poison"] > 0 {
			fmt.Println("2: Vendre 1 Potion de poison (7 Pièces d'or)")
		}
		if p.inventory["Potion de mana"] > 0 {
			fmt.Println("3: Vendre 1 Potion de mana (3 Pièces d'or)")
		}
		if p.inventory["Livre de sort: Boule de feu"] > 0 {
			fmt.Println("4: Vendre 1 Livre de sort: Boule de feu (10 Pièces d'or)")
		}
		fmt.Println("------")
		fmt.Println("5: Page 2, composants")
		fmt.Println("6: Page 3, équipement")
		fmt.Println("0: Retour")
		fmt.Scanln(&input)
		switch input {
		case "1":
			ClearTerminal()
			if p.inventory["Potion de PV"] > 0 {
				p.SellInventory("Potion de PV")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "2":
			ClearTerminal()
			if p.inventory["Potion de poison"] > 0 {
				p.SellInventory("Potion de poison")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "3":
			ClearTerminal()
			if p.inventory["Potion de mana"] > 0 {
				p.SellInventory("Potion de mana")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "4":
			ClearTerminal()
			if p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellInventory("Livre de sort: Boule de feu")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "5":
			ClearTerminal()
			if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
				p.SellDealer2()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer()
			}
		case "6":
			ClearTerminal()
			if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellDealer3()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer()
			}
		case "0":
			ClearTerminal()
			p.Dealer()
		default:
			ClearTerminal()
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.SellDealer()
		}
	} else if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
		p.SellDealer2()
	} else if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
		p.SellDealer3()
	} else {
		fmt.Println("Vous n'avez plus rien à vendre", "\n", "Retour au menu du marchand")
		p.Dealer()
	}
}

func (p *Player) SellDealer2() {
	//Menu vente page 2, seulement si des items sont disponible
	if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
		fmt.Println("Page 2, composants:")
		fmt.Println("Votre or:")
		for i := range p.money {
			fmt.Println(p.money[i], i)
		}
		fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
		fmt.Println("------")
		if p.inventory["Fourrure de loup"] > 0 {
			fmt.Println("1: Vendre 1 Fourrure de loup (2 Pièces d'or)")
		}
		if p.inventory["Peau de troll"] > 0 {
			fmt.Println("2: Vendre 1 Peau de troll (3 Pièces d'or)")
		}
		if p.inventory["Cuir de sanglier"] > 0 {
			fmt.Println("3: Vendre 1 Cuir de sanglier (1 Pièce d'or)")
		}
		if p.inventory["Plume de corbeau"] > 0 {
			fmt.Println("4: Vendre 1 Plume de corbeau ( 0 Pièces d'or)")
		}
		fmt.Println("------")
		fmt.Println("5: Page 3, équipements")
		fmt.Println("0: Page 1, utlitaires")
		fmt.Scanln(&input)
		switch input {
		case "1":
			ClearTerminal()
			if p.inventory["Fourrure de loup"] > 0 {
				p.SellInventory("Fourrure de loup")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "2":
			ClearTerminal()
			if p.inventory["Peau de troll"] > 0 {
				p.SellInventory("Peau de troll")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "3":
			ClearTerminal()
			if p.inventory["Cuir de sanglier"] > 0 {
				p.SellInventory("Cuir de sanglier")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "4":
			ClearTerminal()
			if p.inventory["Plume de corbeau"] > 0 {
				p.SellInventory("Plume de corbeau")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "5":
			ClearTerminal()
			if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellDealer3()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer2()
			}
		case "0":
			ClearTerminal()
			if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellDealer()
			} else {
				fmt.Println("Vous n'avez aucun utilitaires à vendre, retour au menu du marchand")
				p.Dealer()
			}
		default:
			ClearTerminal()
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.SellDealer2()
		}
	} else if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
		p.SellDealer()
	} else if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
		p.SellDealer3()
	} else {
		fmt.Println("Vous n'avez plus rien à vendre", "\n", "Retour au menu du marchand")
		p.Dealer()
	}
}
func (p *Player) SellDealer3() {
	//Menu vente page 3, seulement si les items sont disponible
	if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
		fmt.Println("Page 3, équipements:")
		fmt.Println("Votre or:")
		for i := range p.money {
			fmt.Println(p.money[i], i)
		}
		fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
		fmt.Println("------")
		if p.inventory["Chapeau de l'aventurier"] > 0 {
			fmt.Println("1: Vendre 1 Chapeau de l'aventurier (7 Pièces d'or)")
		}
		if p.inventory["Tunique de l'aventurier"] > 0 {
			fmt.Println("2: Vendre 1 Tunique de l'aventurier (17 Pièces d'or)")
		}
		if p.inventory["Bottes de l'aventurier"] > 0 {
			fmt.Println("3: Vendre 1 Bottes de l'aventurier (10 Pièces d'or)")
		}
		fmt.Println("------")
		fmt.Println("4: Page 2, composants")
		fmt.Println("0: Page 1, utilitaires")
		fmt.Scanln(&input)
		switch input {
		case "1":
			ClearTerminal()
			if p.inventory["Chapeau de l'aventurier"] > 0 {
				p.SellInventory("Chapeau de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "2":
			ClearTerminal()
			if p.inventory["Tunique de l'aventurier"] > 0 {
				p.SellInventory("Tunique de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "3":
			ClearTerminal()
			if p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellInventory("Bottes de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "4":
			ClearTerminal()
			if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
				p.SellDealer2()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer3()
			}
		case "0":
			ClearTerminal()
			if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellDealer()
			} else {
				fmt.Println("Vous n'avez aucun utilitaires à vendre, retour au menu du marchand")
				p.Dealer()
			}
		default:
			ClearTerminal()
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.SellDealer()
		}
	} else if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
		p.SellDealer()
	} else if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
		p.SellDealer2()
	} else {
		fmt.Println("Vous n'avez plus rien à vendre", "\n", "Retour au menu du marchand")
		p.Dealer()
	}
}
