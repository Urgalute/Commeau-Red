package ProjetRed

import (
	"fmt"
	"time"
)

var Turn int

func (p *Player) TrainingFight(turn int) {
	//Menu combat
	fmt.Println("-------------")
	fmt.Println("C'est à votre tour ! (Tour", Turn, ")")
	fmt.Println("---------")
	fmt.Println("1: Attaquer")
	fmt.Println("2: Sort")
	fmt.Println("4: Info ennemie")
	fmt.Println("5: Info personnage")
	fmt.Println("6: Inventaire")
	fmt.Println("0: Retour")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.AttackPlayer(turn)
	case "2":
		p.CastSpell()
	case "4":
		Gobelin.DisplayMonsterInfo()
		p.TrainingFight(turn)
	case "5":
		p.DisplayPlayerInfo()
		p.TrainingFight(turn)
	case "6":
		p.FightInventory()
	case "0":
		p.MainMenu()
	default:
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		fmt.Println("Cette commande ne fait pas partie des possibles, réessayez.")
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		p.TrainingFight(turn)
	}
}

func (p *Player) AttackGoblin(turn int) {
	// Tour ennemie
	time.Sleep(2 * time.Second)
	if Turn == 3 || Turn == 6 || Turn == 9 || Turn == 12 || Turn == 15 || Turn == 18 || Turn == 21 {
		Gobelin.dammage *= 2
	} else {
		Gobelin.dammage = 5
	}
	p.actuallife -= Gobelin.dammage
	if p.actuallife > 0 {
		fmt.Println("L'ennemi vous inflige", Gobelin.dammage, "points de dégâts")
		fmt.Println("Votre vie:", p.actuallife, "/", p.maxlife)
		Turn++
		p.TrainingFight(turn)
	} else {
		p.Wasted()
	}
}

func (p *Player) AttackPlayer(turn int) {
	//Attaque joueur
	Gobelin.actuallife -= p.dammage
	if Gobelin.actuallife > 0 {
		fmt.Println("Vous infligez", p.dammage, "points de dégâts au gobelin")
		fmt.Println("Sa vie:", Gobelin.actuallife, "/", Gobelin.maxlife)
		p.AttackGoblin(turn)
	} else {
		fmt.Println("Vous avez vaincu le gobelin !")
		for i := range Gobelin.loot {
			fmt.Println("Vous gagnez:", Gobelin.loot["Pièces d'or"], i)
			Turn = 1
		}
		p.money["Pièces d'or"] += 5
		fmt.Println("Total:", p.money["Pièces d'or"])
		fmt.Println("Retour au menu principal")
		p.MainMenu()
	}
}
func (p *Player) CastSpell() {
	//Menu sort
	var input string
	fmt.Println("Mana:", p.actualmana, "/", p.maxmana)
	fmt.Println("Sort disponible:")
	fmt.Println("1: Coup de poing, 15 Mana")
	if p.spell[1] == "Boule de feu" {
		fmt.Println("2: Boule de feu, 25 Mana")
	}
	fmt.Println("0: Retour")
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.Punch()
	case "2":
		p.FireBall()
	case "0":
		p.TrainingFight(Turn)
	}
}
func (p *Player) Punch() {
	//Sort coupde poing
	if p.actualmana >= 15 {
		Gobelin.actuallife -= 8
		p.actualmana -= 15
		fmt.Println("Vous infligez 8 points de dégats au gobelin")
		if Gobelin.actuallife >= 0 {
			fmt.Println("Sa vie:", Gobelin.actuallife, "/", Gobelin.maxlife)
			fmt.Println("Votre mana:", p.actualmana, "/", p.maxmana)
			p.AttackGoblin(Turn)
		} else {
			fmt.Println("Vous avez vaincu le gobelin !")
			for i := range Gobelin.loot {
				fmt.Println("Vous gagnez:", Gobelin.loot["Pièces d'or"], i)
				Turn = 1
			}
			p.money["Pièces d'or"] += 5
			fmt.Println("Total:", p.money["Pièces d'or"])
			fmt.Println("Retour au menu principal")
			time.Sleep(3 * time.Second)
			p.MainMenu()
		}
	} else {
		fmt.Println("Vous n'avez pas assez de mana (15):", p.actualmana, "/", p.maxmana)
		p.CastSpell()
	}
}
func (p *Player) FireBall() {
	//Sort boule de feu
	if p.actualmana >= 25 && p.spell[1] == "Boule de feu" {
		Gobelin.actuallife -= 18
		p.actualmana -= 25
		fmt.Println("Vous infligez 18 points de dégats au gobelin")
		if Gobelin.actuallife >= 0 {
			fmt.Println("Sa vie:", Gobelin.actuallife, "/", Gobelin.maxlife)
			fmt.Println("Votre mana:", p.actualmana, "/", p.maxmana)
			p.AttackGoblin(Turn)
		} else {
			fmt.Println("Vous avez vaincu le gobelin !")
			for i := range Gobelin.loot {
				fmt.Println("Vous gagnez:", Gobelin.loot["Pièces d'or"], i)
				Turn = 1
			}
			p.money["Pièces d'or"] += 5
			fmt.Println("Total:", p.money["Pièces d'or"])
			fmt.Println("Retour au menu principal")
			time.Sleep(3 * time.Second)
			p.MainMenu()
		}
	} else if p.actualmana >= 25 && p.spell[1] != "Boule de feu" {
		fmt.Println("Vous n'avez ni assez de mana, ni même appris le sort boule de feu, c'était quoi votre projet ?")
		p.CastSpell()
	} else if p.actualmana < 25 && p.spell[1] == "Boule de feu" {
		fmt.Println("Vous n'avez pas assez de mana (25):", p.actualmana, "/", p.maxmana)
		p.CastSpell()
	} else {
		fmt.Println("Vous n'avez pas encore appris ce sort(Boule de feu)")
		p.TrainingFight(Turn)
	}
}