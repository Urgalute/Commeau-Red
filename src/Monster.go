package ProjetRed

import "fmt"

type Monster struct {
	name       string
	maxlife    int
	actuallife int
	dammage    int
	loot       map[string]int
}

var Gobelin Monster

func InitGoblin() {
	Gobelin.name = "Gobelin d'entrainement"
	Gobelin.maxlife = 40
	Gobelin.actuallife = Gobelin.maxlife
	Gobelin.dammage = 5
	Gobelin.loot = map[string]int{"Pièces d'or": 5}
	Turn = 1
}
func (m Monster) DisplayMonsterInfo() {
	fmt.Println("Nom du monstre:", m.name)
	fmt.Println("Vie max:", m.maxlife)
	fmt.Println("Vie actuelle ", m.actuallife)
	fmt.Println("Dégat", m.dammage)
	for i := range Gobelin.loot {
		fmt.Println("Loot:", Gobelin.loot["Pièces d'or"], i)
	}
	fmt.Println("Cet ennemi inflige 200% de dégât tout les 3 tours")
}
