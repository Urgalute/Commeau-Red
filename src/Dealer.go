package ProjetRed

import (
	"fmt"
)

var input string

func (p *Player) Dealer() {
	fmt.Println("Page 1: Utilitaires")
	fmt.Println("------")
	fmt.Println("Votre or:", p.money)
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("Acheter une potion de PV, 3 Pièces d'or: 1")
	if p.inventory["Potion de PV"] <= 0 {
		fmt.Println("Si vous n'en avez plus, je vous offre la première, si je vous laisse partir sans potion et que vous périssez, cela pourrait nuire à la réputation de mon établissement")
	}
	fmt.Println("Potion de poison, 6 Pièces d'or: 2")
	fmt.Println("Livre de sort: Boule de feu, 25 Pièces d'or : 3")
	if p.rangeinventory == 10 {
		fmt.Println("Capacité d'inventaire +10 (3 Achats restant): 30 Pièces d'or: 4")
	}
	if p.rangeinventory == 20 {
		fmt.Println("Capacité d'inventaire +10 (2 Achats restant): 30 Pièces d'or: 4")
	}
	if p.rangeinventory == 30 {
		fmt.Println("Capacité d'inventaire +10 (1 Achats restant): 30 Pièces d'or: 4")
	}
	if p.rangeinventory == 40 {
		fmt.Println("Capacité d'inventaire +10 (Epuisé, limite d'inventaire max atteinte)")
	}
	fmt.Println("Vendre: 5")
	fmt.Println("Page 2: 6")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case "1":
		if p.inventory["Potion de PV"] < 1 {
			p.AddDealer("Potion de PV", 1, 0)
			p.Dealer()
		} else {
			p.AddDealer("Potion de PV", 1, 3)
			p.Dealer()
		}
	case "2":
		p.AddDealer("Potion de poison", 1, 6)
		p.Dealer()
	case "3":
		p.AddDealer("Livre de sort: Boule de feu", 1, 25)
		p.Dealer()
	case "4":
		if p.rangeinventory <= 30 {
			p.rangeinventory += 10
			fmt.Println("La taille maximale de votre inventaire passe à", p.rangeinventory, "emplacement !")
			p.Dealer()
		} else {
			fmt.Println("Bien essayé mais vous avez déja atteint la capacité d'inventaire max")
			p.Dealer()
		}
	case "5":
		p.SellDealer()
	case "6":
		p.Dealer2()
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer()
	}
}
func (p *Player) Dealer2() {
	fmt.Println("----------------------------")
	fmt.Println("Page 2: Composants")
	fmt.Println("Votre or:", p.money)
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("------")
	fmt.Println("Fourrure de loup, 4 Pièces d'or: 1")
	fmt.Println("Peau de troll, 7 Pièces d'or: 2")
	fmt.Println("Cuir de sanglier, 3 Pièces d'or: 3")
	fmt.Println("Plume de corbeau, 1 Pièces d'or: 4")
	fmt.Println("Retour: 0")
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.AddDealer("Fourrure de loup", 1, 4)
		p.Dealer2()
	case "2":
		p.AddDealer("Peau de troll", 1, 7)
		p.Dealer2()
	case "3":
		p.AddDealer("Cuir de sanglier", 1, 3)
		p.Dealer2()
	case "4":
		p.AddDealer("Plume de corbeau", 1, 1)
		p.Dealer2()
	case "0":
		p.Dealer()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.Dealer2()
	}
}

func (p *Player) SellDealer() {
	var input string
	fmt.Println("Page des ventes:")
	fmt.Println("Votre or:", p.money)
	fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
	fmt.Println("Vos objets:")
	for i := range p.inventory {
		fmt.Println(p.inventory[i], i)
	}
	fmt.Println("------")
	if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
		fmt.Println("Page 1, utilitaires:")
		fmt.Println("Votre or:", p.money)
		fmt.Println("------")
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Vendre 1 Potion de PV (1 Pièce d'or): 1")
		}
		if p.inventory["Potion de poison"] > 0 {
			fmt.Println("Vendre 1 Potion de poison (2 Pièces d'or): 2")
		}
		if p.inventory["Livre de sort: Boule de feu"] > 0 {
			fmt.Println("Vendre 1 Livre de sort: Boule de feu (10 Pièces d'or): 3")
		}
		fmt.Println("------")
		fmt.Println("Page 2, composants: 4")
		fmt.Println("Page 3, équipement: 5")
		fmt.Println("Retour: 0")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.inventory["Potion de PV"] > 0 {
				p.SellInventory("Potion de PV")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "2":
			if p.inventory["Potion de poison"] > 0 {
				p.SellInventory("Potion de poison")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "3":
			if p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellInventory("Livre de sort: Boule de feu")
				p.SellDealer()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer()
			}
		case "4":
			if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
				p.SellDealer2()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer()
			}
		case "5":
			if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellDealer3()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer()
			}
		case "0":
			p.Dealer()
		default:
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
	if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
		fmt.Println("Page 2, composants:")
		fmt.Println("Votre or:", p.money)
		fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
		fmt.Println("------")
		if p.inventory["Fourrure de loup"] > 0 {
			fmt.Println("Vendre 1 Fourrure de loup (2 Pièces d'or): 1")
		}
		if p.inventory["Peau de troll"] > 0 {
			fmt.Println("Vendre 1 Peau de troll (3 Pièces d'or): 2")
		}
		if p.inventory["Cuir de sanglier"] > 0 {
			fmt.Println("Vendre 1 Cuir de sanglier (1 Pièce d'or): 3")
		}
		if p.inventory["Plume de corbeau"] > 0 {
			fmt.Println("Vendre 1 Plume de corbeau ( 0 Pièces d'or): 4")
		}
		fmt.Println("------")
		fmt.Println("Page 3, équipements: 5")
		fmt.Println("Page 1, utlitaires: 0")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.inventory["Fourrure de loup"] > 0 {
				p.SellInventory("Fourrure de loup")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "2":
			if p.inventory["Peau de troll"] > 0 {
				p.SellInventory("Peau de troll")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "3":
			if p.inventory["Cuir de sanglier"] > 0 {
				p.SellInventory("Cuir de sanglier")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "4":
			if p.inventory["Plume de corbeau"] > 0 {
				p.SellInventory("Plume de corbeau")
				p.SellDealer2()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer2()
			}
		case "5":
			if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellDealer3()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer2()
			}
		case "0":
			if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellDealer()
			} else {
				fmt.Println("Vous n'avez aucun utilitaires à vendre, retour au menu du marchand")
				p.Dealer()
			}
		default:
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
	if p.inventory["Chapeau de l'aventurier"] > 0 || p.inventory["Tunique de l'aventurier"] > 0 || p.inventory["Bottes de l'aventurier"] > 0 {
		fmt.Println("Page 3, équipements:")
		fmt.Println("Votre or:", p.money)
		fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
		fmt.Println("------")
		if p.inventory["Chapeau de l'aventurier"] > 0 {
			fmt.Println("Vendre 1 Chapeau de l'aventurier (12 Pièces d'or): 1")
		}
		if p.inventory["Tunique de l'aventurier"] > 0 {
			fmt.Println("Vendre 1 Tunique de l'aventurier (30 Pièces d'or): 2")
		}
		if p.inventory["Bottes de l'aventurier"] > 0 {
			fmt.Println("Vendre 1 Bottes de l'aventurier (18 Pièces d'or): 3")
		}
		fmt.Println("------")
		fmt.Println("Page 2, composants: 4")
		fmt.Println("Page 1, utilitaires: 0")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.inventory["Chapeau de l'aventurier"] > 0 {
				p.SellInventory("Chapeau de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "2":
			if p.inventory["Tunique de l'aventurier"] > 0 {
				p.SellInventory("Tunique de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "3":
			if p.inventory["Bottes de l'aventurier"] > 0 {
				p.SellInventory("Bottes de l'aventurier")
				p.SellDealer3()
			} else {
				fmt.Println("Bien essayé mais vous ne posséder pas cet objet")
				p.SellDealer3()
			}
		case "4":
			if p.inventory["Fourrure de loup"] > 0 || p.inventory["Peau de troll"] > 0 || p.inventory["Cuir de sanglier"] > 0 || p.inventory["Plume de corbeau"] > 0 {
				p.SellDealer2()
			} else {
				fmt.Println("Vous n'avez rien à vendre dans cette catégorie")
				p.SellDealer3()
			}
		case "0":
			if p.inventory["Potion de PV"] > 0 || p.inventory["Potion de poison"] > 0 || p.inventory["Livre de sort: Boule de feu"] > 0 {
				p.SellDealer()
			} else {
				fmt.Println("Vous n'avez aucun utilitaires à vendre, retour au menu du marchand")
				p.Dealer()
			}
		default:
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
