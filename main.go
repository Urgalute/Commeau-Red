package main

import (
	ProjetRed "projetred/src"
)

func main() {
	p1 := ProjetRed.Initplayer("YourName", "Elfe")
	ProjetRed.MainMenu(&p1)
}
