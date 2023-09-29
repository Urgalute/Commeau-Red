package ProjetRed

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"os"
	"time"
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
		ClearTerminal()
		fmt.Println("----------------------------")
		p.DisplayPlayerInfo()
		p.MainMenu()
	case "2":
		ClearTerminal()
		fmt.Println("----------------------------")
		p.AccessInventory()
	case "3":
		ClearTerminal()
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le marchand !", "\n", "Je vous propose quelques objet pour votre aventure, je peux également vous racheter vos objets")
		fmt.Println("------")
		p.Dealer()
	case "4":
		ClearTerminal()
		fmt.Println("----------------------------")
		fmt.Println("Bienvenue chez le forgeron !", "\n", "Contre une légère commission, ma forge est à votre disposition.")
		fmt.Println("------")
		p.BlackSmith()
	case "5":
		ClearTerminal()
		Randinitgob = rand.Intn(20-1) + 1
		Randinitplayer = rand.Intn(20-1) + 1
		fmt.Println("Vous rencontrez un Gobelin !")
		InitGoblin()
		fmt.Println("Jet d'initiative (Jet sur 20), vous avez fait:", Randinitplayer, "Le gobelin a fait:", Randinitgob)
		if Randinitgob > Randinitplayer {
			fmt.Println("Le gobelin commence !")
			time.Sleep(2 * time.Second)
			p.AttackGoblin(Turn)
		} else {
			fmt.Println("Vous commencez !")
			p.TrainingFight(Turn)
		}
	case "9":
		ClearTerminal()
		p.Whoarethey()
	case "0":
		ClearTerminal()
		p.Sure()
	default:
		ClearTerminal()
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


func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
}
}
func (p *Player) Sure(){
fmt.Println("Vous êtes sur le point de nous quitter, êtes-vous vraiment sur ?")
fmt.Println("5: Certain")
fmt.Println("0: Non, je reste")
fmt.Scanln(&input)
switch input {

case "5": 
time.Sleep(2 * time.Second)
fmt.Println("A la prochaine !!")
case "0": 
time.Sleep(2 * time.Second)
p.MainMenu()
}
}