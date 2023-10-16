package game

import (
	"github.com/fatih/color"
	"projet/master"
	"projet/utils"
)

// Affiche l'interface de sélection du dictionnaire
func SelectDictionnary() {
	utils.ClearTerminal()
	color.Cyan("Quel est le nom du fichier :")
	color.Yellow(master.WritingWord)
	color.Cyan("Appuyer sur entrée pour confirmer")
}

// Gère lorsqu'on appuie sur "Entrée" dans le menu du dictionnaire
func ExecForDictionnary(r rune) {
	if r == 13 { // touche entrée
		master.IsWriting = false
		master.TextBook = master.WritingWord
		master.WritingWord = ""
		master.CurrentPage = "menu"
		DisplayMenu()
	} else if r >= 97 && r <= 122 || r == 46 || r == 95 {
		master.WritingWord += string(rune(r))
		SelectDictionnary()
	}
}
