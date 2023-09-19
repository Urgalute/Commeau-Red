package Red

import "fmt"

type player struct {
	name      string
	classe    string
	level     int
	pvmax     int
	pvactual  int
	inventory map[string]int
}

func (p player) InitPlayer (a string, b string){
	var p1 player

	p1.name = a
	p1.classe = b
	p1.level = 30
	p1.pvmax = 800
	p1.pvactual = 500
	p1.inventory = map[string]int{"Potion": 3, "Epée": 1}

	p1.DisplayInfo()
	p1.AccessInventory()
	p1.TakePot()
	p1.AccessInventory()
	p1.DisplayInfo()
}

func (p player) DisplayInfo() {
	fmt.Println("Nom =", p.name)
	fmt.Println("Classe =", p.classe)
	fmt.Println("Niveau =", p.level)
	fmt.Println("PV max. =", p.pvmax)
	fmt.Println("PV =", p.pvactual)
}
