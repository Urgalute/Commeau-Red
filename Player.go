package ProjetRed

import "fmt"

type player struct {
	name       string
	class      string
	level      int
	maxlife    int
	actuallife int
	inventory  map[string]int
}

func Initplayer(name string, class string) {

	var p1 player
	p1.name = name
	p1.class = class
	p1.level = 3
	p1.maxlife = 200
	p1.actuallife = 140
	p1.inventory = map[string]int{"Potion": 3, "Epée": 1}

	p1.DisplayPlayerInfo()
	p1.AccessInventory()
	p1.TakePot()
	fmt.Println("------------------")
	p1.DisplayPlayerInfo()
	p1.AccessInventory()
}

func (p player) DisplayPlayerInfo() {
	fmt.Println("Nom du personnage:", p.name)
	fmt.Println("Classe: ", p.class)
	fmt.Println("Niveau du personnage: ", p.level)
	fmt.Println("Maximum de PV: ", p.maxlife)
	fmt.Println("PV actuel: ", p.actuallife)
}
