package game

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"projet/master"
	"projet/utils"
)

// Permet d'afficher l'interface du menu principal
func DisplayMenu() {
	utils.ClearTerminal()
	switch master.MenuIndex {
	case 1:
		data := [][]string{
			[]string{"> Jouer"}, []string{"Changer de dictionnaire | " + master.TextBook}, []string{"Quitter"},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		for _, v := range data {
			table.Append(v)
		}
		table.SetFooter([]string{"Espace pour intéragir"})
		table.Render()
		break
	case 2:
		data := [][]string{
			[]string{"Jouer"}, []string{"> Changer de dictionnaire | " + master.TextBook}, []string{"Quitter"},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		for _, v := range data {
			table.Append(v)
		}
		table.SetFooter([]string{"Espace pour intéragir"})
		table.Render()
		break
	case 3:
		data := [][]string{
			[]string{"Jouer"}, []string{"Changer de dictionnaire | " + master.TextBook}, []string{"> Quitter"},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		for _, v := range data {
			table.Append(v)
		}
		table.SetFooter([]string{"Espace pour intéragir"})
		table.Render()
		break
	}
}

// Gère les actions effectuées lorsqu'on appuie sur "Entrée"
func ExecMenu(r rune) {
	if r == 65 {
		if master.MenuIndex == 1 {
			master.MenuIndex = 3
		} else {
			master.MenuIndex--
		}
		DisplayMenu()
	} else if r == 66 {
		if master.MenuIndex == 3 {
			master.MenuIndex = 1
		} else {
			master.MenuIndex++
		}
		DisplayMenu()
	}
	if r == 32 {
		switch master.MenuIndex {
		case 1:
			master.MenuIndex = 0
			master.CurrentPage = "ingame"
			NewGame()
			break
		case 2:
			master.IsWriting = true
			master.WritingWord = ""
			master.CurrentPage = "select"
			SelectDictionnary()
			break
		case 3:
			os.Exit(0)
		}
	}
}
