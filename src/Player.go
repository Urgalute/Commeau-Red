package ProjetRed

import "fmt"

type Equipment struct {
	head  map[string]int
	torse map[string]int
	foot  map[string]int
}
type Player struct {
	name           string
	race           string
	xp int
	xpmax int	
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
		p1.maxlife = 100
		p1.maxmana = 50
	case "Elfe":
		p1.maxlife = 80
		p1.maxmana = 75
	case "Nain":
		p1.maxlife = 120
		p1.maxmana = 25
	}
	p1.xp = 0
	p1.xpmax = 20
	p1.level = 1
	p1.dammage = 5
	p1.actuallife = p1.maxlife / 2
	p1.actualmana = p1.maxmana / 2
	p1.spell = [2]string{"Coup de poing"}
	p1.money = map[string]int{"Pièces d'or": 50}
	p1.inventory = map[string]int{"Potion de PV": 3}
	p1.rangeinventory = 5
	p1.equipment.head = map[string]int{}
	p1.equipment.torse = map[string]int{}
	p1.equipment.foot = map[string]int{}

	Gobelin.maxlife = 40
	Gobelin.dammage = 3
	Gobelin.level = 1
	Gobelin.xp = 5
	Gobelin.loot = map[string]int{"Pièces d'or": 10}

	p1.DisplayPlayerInfo()
	p1.MainMenu()
	return p1
}
func (p *Player) DisplayPlayerInfo() {
	fmt.Println("Nom du personnage:", p.name)
	fmt.Println("Race:", p.race)
	if p.level != 6 {
	fmt.Println("Niveau du personnage: ", p.level)
	fmt.Println("Mon expérience:", p.xp,"/", p.xpmax)} else {
		fmt.Println("Niveau du personnage: ", p.level)
	}
	fmt.Println("Dégâts:", p.dammage)
	fmt.Println("PV actuel / Maximum de PV: ", p.actuallife, "/", p.maxlife)
	fmt.Println("Mana actuel / Maximum de mana:", p.actualmana, "/", p.maxmana)
	fmt.Println("Sorts:")
	for i := range p.spell {
		fmt.Println(p.spell[i])
	}
	fmt.Println("Equipement:")
	for i := range p.equipment.head {
		if p.equipment.head[i] < 0 {
			fmt.Println(i, "0 / 50")
		} else {
			fmt.Println(i, p.equipment.head[i], "/ 50")
		}
	}
	for i := range p.equipment.torse {
		if p.equipment.torse[i] < 0 {
			fmt.Println(i, "0 / 50")
		} else {
			fmt.Println(i, p.equipment.torse[i], "/ 50")
		}
	}
	for i := range p.equipment.foot {
		if p.equipment.foot[i] < 0 {
			fmt.Println(i, "0 / 50")
		} else {
			fmt.Println(i, p.equipment.foot[i], "/ 50")
		}
	}
	fmt.Println("----------------------------")
}
func (p *Player) Experience(){
	//var p1 Player
	if p.level == 6 {
		fmt.Println("Vous avez atteint le niveau maximum")
	} else {
		p.xp += Gobelin.xp 
		fmt.Println("Vous gagnez", Gobelin.xp,"points d'expériences")
	switch p.xp{
		case 20: p.level = 2 ; p.xpmax = 40 
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 40: p.level = 3 ; p.xpmax = 60
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 60: p.level = 4 ; p.xpmax = 80
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 80: p.level = 5 ; p.xpmax = 100
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 100: p.level = 6
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,"(Niveau maximum). Vos statistique augmentent, mais celles de vos ennemies aussi !")
		}	
	switch p.level{
		case 1: p.dammage = 5 ; Gobelin.xp = 5 ; Gobelin.maxlife = 40 ; Gobelin.dammage = 3 ; Gobelin.level = 1
		case 2: p.dammage = 8 ; Gobelin.xp = 10 ; Gobelin.maxlife = 50 ; Gobelin.dammage = 5 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 2
		case 3: p.dammage = 11 ; Gobelin.xp = 15 ; Gobelin.maxlife = 60 ; Gobelin.dammage = 7 ; p.maxlife += 5 ; p.maxmana += 10  ; Gobelin.level = 3; Gobelin.loot["Pièces d'or"] = 13
		case 4: p.dammage = 14 ; Gobelin.xp = 20 ; Gobelin.maxlife = 70 ; Gobelin.dammage = 9 ; p.maxlife += 5 ; p.maxmana +=10  ; Gobelin.level = 4
		case 5: p.dammage = 17 ; Gobelin.xp = 25 ; Gobelin.maxlife = 80 ; Gobelin.dammage = 11 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 5
		case 6: p.dammage = 20 ; Gobelin.maxlife = 90 ; Gobelin.dammage = 13 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 6 ; Gobelin.loot["Pièces d'or"] = 20
		}
	}
}