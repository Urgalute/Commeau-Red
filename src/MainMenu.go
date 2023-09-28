package ProjetRed

import (
	"fmt"
	"math/rand"
)

var Randinitgob int
var Randinitplayer int

func (p *Player) MainMenu() {
	var input string
	fmt.Println("----------------------------")
	fmt.Println("Menu principal:")
	fmt.Println("------")
	fmt.Println("1: Information personnage")
	fmt.Println("2: Inventaire")
	fmt.Println("3: Marchand")
	fmt.Println("4: Forgeron")
	fmt.Println("5: Combat d'entrainement")
	fmt.Println("9: WhoAreThey?")
	fmt.Println("0: Quitter")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("----------------------------")
		p.DisplayPlayerInfo()
		p.MainMenu()
	case "2":
		fmt.Println("----------------------------")
		p.AccessInventory()
	case "3":
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le marchand !", "\n", "Je vous propose quelques objet pour votre aventure, je peux également vous racheter vos objets")
		fmt.Println("------")
		p.Dealer()
	case "4":
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le forgeron !", "\n", "Contre une légère commission, ma forge est à votre disposition.")
		fmt.Println("------")
		p.BlackSmith()
	case "5":
		Randinitgob = rand.Intn(20-1) + 1
		Randinitplayer = rand.Intn(20-1) + 1
		fmt.Println("Vous rencontrez un Gobelin !")
		InitGoblin()
		fmt.Println("Jet d'initiative:", Randinitplayer, "Le gobelin a fait:", Randinitgob)
		if Randinitgob > Randinitplayer {
			fmt.Println("Le gobelin commence !")
			p.AttackGoblin(Turn)
		} else {
			fmt.Println("Vous commencez !")
			p.TrainingFight(Turn)
		}
	case "9":
		p.Whoarethey()
	case "0":
		fmt.Println("----------------------------")
		fmt.Println("A plus !")
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.MainMenu()
	}
}
func (p *Player) Whoarethey() {
	fmt.Println("ABBA, Steven Spielberg, Queen")
	p.MainMenu()
}
