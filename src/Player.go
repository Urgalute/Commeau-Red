package ProjetRed

import "fmt"

type Equipment struct {
	head  string
	torse string
	foot  string
}
type Player struct {
	name           string
	race           string
	level          int
	maxlife        int
	actuallife     int
	actualmana     int
	maxmana        int
	spell          [2]string
	dammage        int
	equipment      Equipment
	money          map[string]int
	inventory      map[string]int
	rangeinventory int
}

func Initplayer(name string, race string) Player {

	var p1 Player
	p1.name = name
	p1.race = race
	switch p1.race {
	default:
		p1.maxlife = 100
	case "Humain":
		p1.maxlife = 125
	case "Elfe":
		p1.maxlife = 80
	case "Nain":
		p1.maxlife = 120
	}
	p1.level = 1
	p1.actuallife = p1.maxlife / 2
	p1.dammage = 10
	p1.maxmana = 50
	p1.actualmana = p1.maxmana / 2
	p1.spell = [2]string{"Coup de poing"}
	p1.money = map[string]int{"Pièces d'or": 200}
	p1.inventory = map[string]int{"Potion de PV": 3}
	p1.rangeinventory = 10
	p1.equipment.head = ""
	p1.equipment.torse = ""
	p1.equipment.foot = ""
	p1.DisplayPlayerInfo()
	p1.MainMenu()
	return p1
}
func (p *Player) DisplayPlayerInfo() {
	fmt.Println("Nom du personnage:", p.name)
	fmt.Println("Race:", p.race)
	fmt.Println("Niveau du personnage: ", p.level)
	fmt.Println("PV actuel / Maximum de PV: ", p.actuallife, "/", p.maxlife)
	fmt.Println("Mana actuel / Maximum de mana:", p.actualmana, "/", p.maxmana)
	fmt.Println("Sorts:")
	for i := range p.spell {
		fmt.Println(p.spell[i])
	}
	fmt.Println("Equipement:", p.equipment)
	fmt.Println("----------------------------")
	
}
