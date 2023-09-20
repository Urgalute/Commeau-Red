package ProjetRed

import "fmt"

func (p *Player) BlackSmith() {
	if p.money >= 5 {
		fmt.Println("Voilà ce que je peut vous proposer.")
		{
			if p.inventory["Plume de corbeau"] >= 1 && p.inventory["Cuir de sanglier"] >= 1 {
				fmt.Println("Chapeau de l'aventurier (5 Pièces d'or, 1 Plume de corbeau, 1 Cuir de sanglier): 1")
			} else {
				fmt.Println("Il vous manque quelques composants pour fabriquer le chapeau de l'aventurier:", "Plume de corbeau:", p.inventory["Plume de corbeau"], "/", "1,", "Cuir de sanglier", p.inventory["Cuir de sanglier"], "/", "1")
			}
			if p.inventory["Fourrure de loup"] >= 2 && p.inventory["Peau de troll"] >= 1 {
				fmt.Println("Tunique de l'aventurier (5 Pièces d'or, 2 Fourrure de loup, 1 Peau de troll): 2")
			} else {
				fmt.Println("Il vous manque quelques composants pour fabriquer la tunique de l'aventurier:", "Fourrure de loup:", p.inventory["Fourrure de loup"], "/", "2,", "Peau de troll", p.inventory["Peau de troll"], "/", "1")
			}
			if p.inventory["Fourrure de loup"] >= 1 && p.inventory["Cuir de sanglier"] >= 1 {
				fmt.Println("Bottes de l'aventurier (5 Pièces d'or, 1 Fourrure de loup, 1 Cuir de sanglier): 3")
			} else {
				fmt.Println("Il vous manque quelques composants pour fabriquer les bottes de l'aventurier:", "Fourrure de loup:", p.inventory["Fourrure de loup"], "/", "1,", "Cuir de sanglier", p.inventory["Cuir de sanglier"], "/", "1")
			}
		}

		fmt.Println("Retour: 0")
		fmt.Scanln(&input)
		switch input {
		case 1:
			if p.inventory["Plume de corbeau"] > 0 && p.inventory["Cuir de sanglier"] > 0 {
				p.AddCraft("Chapeau de l'aventurier", 1, 5)
				p.inventory["Plume de corbeau"]--
				p.inventory["Cuir de sanglier"]--
				p.ChechInventory()
				p.BlackSmith()
			}
		case 2:
			if p.inventory["Fourrure de loup"] > 1 && p.inventory["Peau de troll"] > 0 {
				p.AddCraft("Tunique de l'aventurier", 1, 5)
				p.inventory["Fourrure de loup"] -= 2
				p.inventory["Peau de troll"]--
				p.ChechInventory()
				p.BlackSmith()
			}
		case 3:
			if p.inventory["Fourrure de loup"] > 0 && p.inventory["Cuir de sanglier"] > 0 {
				p.AddCraft("Bottes de l'aventurier", 1, 5)
				p.inventory["Fourrure de loup"]--
				p.inventory["Cuir de sanglier"]--
				p.ChechInventory()
				p.BlackSmith()
			}
		case 0:
			p.MainMenu()
		default:
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
			fmt.Println("---------------------------------------------------------------------------------------------------------")
			p.BlackSmith()
		}
	} else {
		fmt.Println("Désolé mais je ne travaille pas gratuitement..", "\n", "(Il vous faut 5 Pièces d'or au minimum pour demander au forgeron de vous fabriquer un objet. Vous avez:", p.money, ")")
		fmt.Println("Retour au menu principal")
		p.MainMenu()
	}
}
