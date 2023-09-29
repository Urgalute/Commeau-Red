package ProjetRed

import "fmt"

type Monster struct {
	name       string
	maxlife    int
	actuallife int
	dammage    int
	xp int
	level int
	loot       map[string]int
}

var Gobelin Monster

func InitGoblin() {
	Gobelin.name = "Gobelin d'entrainement"
	Gobelin.actuallife = Gobelin.maxlife

	Turn = 1

}
func (m Monster) DisplayMonsterInfo() {
	fmt.Println("Nom du monstre:", Gobelin.name)
fmt.Println("Niveau:", Gobelin.level)
	fmt.Println("Vie max:", Gobelin.maxlife)
	fmt.Println("Vie actuelle ", Gobelin.actuallife)
	fmt.Println("Dégat", Gobelin.dammage)
	for i := range Gobelin.loot {
		fmt.Println("Loot:", Gobelin.loot["Pièces d'or"], i)
	}
	fmt.Println("Cet ennemi inflige 200% de dégât tout les 3 tours")
}
