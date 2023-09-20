package ProjetRed

import "fmt"

type Player struct {
	name       string
	race       string
	class      string
	level      int
	maxlife    int
	actuallife int
	money      int
	inventory  map[string]int
}

func Initplayer(name string, race string, class string) Player {

	var p1 Player
	p1.name = name
	p1.race = race
	p1.class = class
	p1.level = 1
	p1.maxlife = 200
	p1.actuallife = 100
	p1.money = 100
	p1.inventory = map[string]int{"Potion de PV": 3}
	p1.DisplayPlayerInfo()
	return p1
}
func (p *Player) DisplayPlayerInfo() {
	fmt.Println("Nom du personnage:", p.name)
	fmt.Println("Race:", p.race)
	fmt.Println("Classe: ", p.class)
	fmt.Println("Niveau du personnage: ", p.level)
	fmt.Println("Maximum de PV: ", p.maxlife)
	fmt.Println("PV actuel: ", p.actuallife)
	fmt.Println("----------------------------")
	fmt.Println("Retour: 0")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 0:
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas parite des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.DisplayPlayerInfo()
	}
}
