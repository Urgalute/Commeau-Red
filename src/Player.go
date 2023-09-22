package ProjetRed

import "fmt"

type Equipment struct {
	head  string
	torse string
	foot  string
}
type Spell struct {
	spell map[string]int
}
type Player struct {
	name       string
	race       string
	class      string
	level      int
	maxlife    int
	actuallife int
	equipment  Equipment
	spell      map[string]int
	money      int
	inventory  map[string]int
	rangeinventory int
}

func Initplayer(name string ,race string, class string) Player {

	var p1 Player
	p1.name = name
	p1.race = race
	switch p1.race {
	default:
		p1.maxlife = 100
	case "Humain":
		p1.maxlife = 100
	case "Elfe":
		p1.maxlife = 80
	case "Nain":
		p1.maxlife = 120
	}
	p1.class = class
	p1.level = 1
	p1.actuallife = p1.maxlife / 2
	p1.spell = map[string]int{"Coup de poing": 1}
	p1.money = 300
	p1.inventory = map[string]int{"Potion de PV": 1}
	p1.rangeinventory = 10
	p1.equipment.head = ""
	p1.equipment.torse = ""
	p1.equipment.foot = ""
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
	fmt.Println("Equipement:", p.equipment)
	fmt.Println("----------------------------")
	fmt.Println("Retour: 0")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.DisplayPlayerInfo()
	}
}
