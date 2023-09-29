package ProjetRed

import (
	"fmt"
	"time"
)

// Prendre une potion de PV
func (p *Player) TakePot() {
	if p.actuallife > p.maxlife-50 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion de vie")
	}
	if p.actuallife <= p.maxlife-50 && p.inventory["Potion de PV"] > 0 {
		p.inventory["Potion de PV"]--
		p.actuallife += 50
		fmt.Println("----------------------------")
		fmt.Println("Une potion de PV a été utilisé, nouveau montant de PV:", p.actuallife, "/", p.maxlife)
		p.DeleteInventory()
		if p.inventory["Potion de PV"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de PV"])
		}
		fmt.Println("----------------------------")
	}
}

// Potion de poison
func (p *Player) PoisonPot() {
	if p.inventory["Potion de poison"] > 0 {
		fmt.Println("A l'avenir quand vous verrez un liquide vert, fûmant, dans une fiole en verre, évitez de le boire..")
		time.Sleep(3 * time.Second)
		for i := 0; i < 3; i++ {
			p.actuallife -= 10
			fmt.Println("Votre vie:", p.actuallife, "/", p.maxlife)
			time.Sleep(1 * time.Second)
		}
		p.inventory["Potion de poison"]--
		p.DeleteInventory()
		p.Wasted()
	}
}

// Livre de sort: Boule de feu
func (p *Player) SpellBook() {
	if p.inventory["Livre de sort: Boule de feu"] > 0 {
		p.spell[1] = "Boule de feu"
		p.inventory["Livre de sort: Boule de feu"]--
		p.DeleteInventory()
		fmt.Println("Vous avez appris un nouveau sort: Boule de feu")
	}
}

// Potion de mana
func (p *Player) TakeManaPot() {
	if p.actualmana > p.maxmana-25 {
		fmt.Println("----------------------------")
		fmt.Println("Vous ne pouvez pas prendre de potion de mana")
	}
	if p.actualmana <= p.maxmana-25 && p.inventory["Potion de mana"] > 0 {
		p.inventory["Potion de mana"]--
		p.actualmana += 25
		fmt.Println("----------------------------")
		fmt.Println("Une potion de mana a été utilisé, nouveau montant de mana:", p.actualmana, "/", p.maxmana)
		p.DeleteInventory()
		if p.inventory["Potion de mana"] > 0 {
			fmt.Println("Il vous en reste:", p.inventory["Potion de mana"])
		}
		fmt.Println("----------------------------")
	}
}
func (p *Player) TakeHammer(){
	var i int
	if p.equipment.head["Chapeau de l'aventurier"] == 50 {
   fmt.Println("La durabilité de votre Chapeau de l'aventurier est déjà au maximum")
} else if p.equipment.head["Chapeau de l'aventurier"] == -5 || p.equipment.head["Chapeau de l'aventurier"] > 0 && p.equipment.head["Chapeau de l'aventurier"] != 50 {
   if p.equipment.head["Chapeau de l'aventurier"] == -5 {
	   p.maxlife += 10
	   p.maxmana += 10
	   fmt.Println("Vous recupérez votre bonus de caractéristique")
   }
   p.equipment.head["Chapeau de l'aventurier"] = 50
   fmt.Println("Vous avez bien réparé votre Chapeau de l'aventurier")
   i = 1
} else {
   fmt.Println("Vous n'avez rien d'équipé sur votre tête")
}
if p.equipment.torse["Tunique de l'aventurier"] == 50 {
   fmt.Println("La durabilité de votre Tunique de l'aventurier est déjà au maximum")
} else if p.equipment.torse["Tunique de l'aventurier"] == -5 || p.equipment.torse["Tunique de l'aventurier"] > 0 && p.equipment.torse["Tunique de l'aventurier"] != 50 {
   if p.equipment.torse["Tunique de l'aventurier"] == -5 {
	   p.maxlife += 25
	   p.maxmana += 25
	   fmt.Println("Vous recupérez votre bonus de caractéristique")
   }
   p.equipment.torse["Tunique de l'aventurier"] = 50
   fmt.Println("Vous avez bien réparé votre Tunique de l'aventurier")
   i = 1
} else {
   fmt.Println("Vous n'avez rien d'équipé sur votre torse")
}
if p.equipment.foot["Bottes de l'aventurier"] == 50 {
   fmt.Println("La durabilité de vos Bottes de l'aventurier est déjà au maximum")
} else if p.equipment.foot["Bottes de l'aventurier"] == -5 || p.equipment.foot["Bottes de l'aventurier"] > 0 && p.equipment.foot["Bottes de l'aventurier"] != 50 {
   if p.equipment.foot["Bottes de l'aventurier"] == -5 {
	   p.maxlife += 15
	   p.maxmana += 15
	   fmt.Println("Vous recupérez votre bonus de caractéristique")
   }
   p.equipment.foot["Bottes de l'aventurier"] = 50
   fmt.Println("Vous avez bien réparé vos Bottes de l'aventurier")
   i = 1
} else {
   fmt.Println("Vous n'avez rien d'équipé sur vos jambes")
}
if i == 1{
p.inventory["Marteau de forgeron"]--
p.DeleteInventory()}
}

func (p *Player) TakeHammer2(){
	var i int 
	if p.equipment.head["Chapeau de l'aventurier"] == 50 {
		fmt.Println("La durabilité de votre Chapeau de l'aventurier est déjà au maximum")
	 } else if p.equipment.head["Chapeau de l'aventurier"] == -5 || p.equipment.head["Chapeau de l'aventurier"] > 0 && p.equipment.head["Chapeau de l'aventurier"] != 50 {
		if p.equipment.head["Chapeau de l'aventurier"] == -5 {
			p.maxlife += 10
			p.maxmana += 10
			fmt.Println("Vous recupérez votre bonus de caractéristique")
		}
		p.equipment.head["Chapeau de l'aventurier"] = 50
		fmt.Println("Vous avez bien réparé votre Chapeau de l'aventurier")
		i = 1
	 } else {
		fmt.Println("Vous n'avez rien d'équipé sur votre tête")
	 }
	 if p.equipment.torse["Tunique de l'aventurier"] == 50 {
		fmt.Println("La durabilité de votre Tunique de l'aventurier est déjà au maximum")
	 } else if p.equipment.torse["Tunique de l'aventurier"] == -5 || p.equipment.torse["Tunique de l'aventurier"] > 0 && p.equipment.torse["Tunique de l'aventurier"] != 50 {
		if p.equipment.torse["Tunique de l'aventurier"] == -5 {
			p.maxlife += 25
			p.maxmana += 25
			fmt.Println("Vous recupérez votre bonus de caractéristique")
		}
		p.equipment.torse["Tunique de l'aventurier"] = 50
		fmt.Println("Vous avez bien réparé votre Tunique de l'aventurier")
		i = 1
	 } else {
		fmt.Println("Vous n'avez rien d'équipé sur votre torse")
	 }
	 if p.equipment.foot["Bottes de l'aventurier"] == 50 {
		fmt.Println("La durabilité de vos Bottes de l'aventurier est déjà au maximum")
	 } else if p.equipment.foot["Bottes de l'aventurier"] == -5 || p.equipment.foot["Bottes de l'aventurier"] > 0 && p.equipment.foot["Bottes de l'aventurier"] != 50 {
		if p.equipment.foot["Bottes de l'aventurier"] == -5 {
			p.maxlife += 15
			p.maxmana += 15
			fmt.Println("Vous recupérez votre bonus de caractéristique")
		}
		p.equipment.foot["Bottes de l'aventurier"] = 50
		fmt.Println("Vous avez bien réparé vos Bottes de l'aventurier")
		i = 1
	 } else {
		fmt.Println("Vous n'avez rien d'équipé sur vos jambes")
	 }
	 if i == 1{
	p.money["Pièces d'or"] -=10
	}
	 }
