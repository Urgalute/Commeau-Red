package ProjetRed

import "fmt"

func (p *Player) BlackSmith() {
	if p.money["Pièces d'or"] >= 5 {
		fmt.Println("Votre or:")
		for i := range p.money {
			fmt.Println(p.money[i], i)
		}
		fmt.Println("Capacité de l'inventaire:", len(p.inventory), "/", p.rangeinventory)
		fmt.Println("Voilà ce que je peut vous proposer.")
		{
			if p.inventory["Plume de corbeau"] >= 1 && p.inventory["Cuir de sanglier"] >= 1 {
				fmt.Println("1: Chapeau de l'aventurier (5 Pièces d'or, 1 Plume de corbeau, 1 Cuir de sanglier)")
			} else {
				fmt.Println("Chapeau de l'aventurie: Il vous manque quelques composants pour le fabriquer:", "Plume de corbeau:", p.inventory["Plume de corbeau"], "/", "1,", "Cuir de sanglier", p.inventory["Cuir de sanglier"], "/", "1")
			}
			if p.inventory["Fourrure de loup"] >= 2 && p.inventory["Peau de troll"] >= 1 {
				fmt.Println("2: Tunique de l'aventurier (5 Pièces d'or, 2 Fourrure de loup, 1 Peau de troll)")
			} else {
				fmt.Println("Tunique de l'aventurier: Il vous manque quelques composants pour la fabriquer:", "Fourrure de loup:", p.inventory["Fourrure de loup"], "/", "2,", "Peau de troll", p.inventory["Peau de troll"], "/", "1")
			}
			if p.inventory["Fourrure de loup"] >= 1 && p.inventory["Cuir de sanglier"] >= 1 {
				fmt.Println("3: Bottes de l'aventurier (5 Pièces d'or, 1 Fourrure de loup, 1 Cuir de sanglier)")
			} else {
				fmt.Println("Bottes de l'aventurier: Il vous manque quelques composants pour les fabriquer:", "Fourrure de loup:", p.inventory["Fourrure de loup"], "/", "1,", "Cuir de sanglier", p.inventory["Cuir de sanglier"], "/", "1")
			}
		}
		var input string
		fmt.Println("0: Retour")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.inventory["Plume de corbeau"] > 0 && p.inventory["Cuir de sanglier"] > 0 {
				p.AddCraft("Chapeau de l'aventurier", 1, 5)
				p.inventory["Plume de corbeau"]--
				p.inventory["Cuir de sanglier"]--
				p.DeleteInventory()
				p.BlackSmith()
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
				p.BlackSmith()
			}
		case "2":
			if p.inventory["Fourrure de loup"] > 1 && p.inventory["Peau de troll"] > 0 {
				p.AddCraft("Tunique de l'aventurier", 1, 5)
				p.inventory["Fourrure de loup"] -= 2
				p.inventory["Peau de troll"]--
				p.DeleteInventory()
				p.BlackSmith()
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
				p.BlackSmith()
			}
		case "3":
			if p.inventory["Fourrure de loup"] > 0 && p.inventory["Cuir de sanglier"] > 0 {
				p.AddCraft("Bottes de l'aventurier", 1, 5)
				p.inventory["Fourrure de loup"]--
				p.inventory["Cuir de sanglier"]--
				p.DeleteInventory()
				p.BlackSmith()
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
				p.BlackSmith()
			}
		case "0":
			p.MainMenu()
		default:
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.BlackSmith()
		}
	} else {
		fmt.Println("Désolé mais je ne travaille pas gratuitement..", "\n", "(Il vous faut 5 Pièces d'or au minimum pour demander au forgeron de vous fabriquer un objet. Vous avez:", p.money, ")")
		fmt.Println("Retour au menu principal")
		p.MainMenu()
	}
}
