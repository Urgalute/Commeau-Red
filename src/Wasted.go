package ProjetRed

import "fmt"

func (p *Player) Wasted() bool {
	i := false
	if p.actuallife <= 0 {
		p.actuallife = p.maxlife / 2
		fmt.Println("Vous êtes mort..", "\n", "Vous revenez au menu principal avec", p.actuallife, "PV", "(50% de vos PV max)")
		i = true
		p.MainMenu()
	}
	return i
}
