package main

import (
	"github.com/mattn/go-tty"
	"log"
	"os"
	"projet/game"
	"projet/master"
)

func main() {
	// change la cible du .txt
	if len(os.Args) >= 2 && os.Args[1] != "" {
		master.TextBook = os.Args[1]
	}
	game.DisplayMenu()
	StartListening()
}

// Gère l'écoute des inputs
func StartListening() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if !master.CanPlay && master.CurrentPage == "menu" {
			game.ExecMenu(r)
		} else if !master.CanPlay && master.CurrentPage == "select" {
			game.ExecForDictionnary(r)
		}
		if r >= 97 && r <= 122 || r == 45 || r == 13 || r == 8 {
			// si les touches sont des lettres, tiret, entrée
			if !master.CanPlay && master.CurrentPage == "ingame" {
				game.ReturnIntoMainMenu()
			} else if master.CanPlay {
				// la partie est en cours, on capture donc les inputs
				if r == 13 { // touche entrée
					game.ExecIntoGame(r)
				} else {
					if master.IsWriting {
						// on capture les touches pour un mot entier donc on ajoute la lettre
						game.WritingWord(r)
					} else {
						// on test la lettre si il nous reste des tentatives
						if master.Chances > 0 {
							game.TestLetter(r)
						} else {
							game.NoMoreChance()
						}
					}
				}
			}
		}
	}
}
