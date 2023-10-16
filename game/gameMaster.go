package game

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"math/rand"
	"os"
	"projet/master"
	"projet/utils"
	"strconv"
	"strings"
)

// Initialise une nouvelle partie
func NewGame() {
	master.CanPlay = true
	master.Chances = 10
	master.Errors = 0
	master.WordToFind = utils.GetRandomWord()
	master.WordToFind = strings.ReplaceAll(master.WordToFind, "\n", "")
	master.HiddenWord = strings.Repeat("_", len(master.WordToFind))
	master.GivenLetters = make(map[rune]bool)
	master.UsedLetters = []string{}
	letter1 := rand.Intn(len(master.WordToFind))
	letter2 := rand.Intn(len(master.WordToFind))
	master.GivenLetters[rune(master.WordToFind[letter1])] = true
	master.GivenLetters[rune(master.WordToFind[letter2])] = true
	master.UsedLetters = append(master.UsedLetters, string(rune(master.WordToFind[letter1])))
	if rune(master.WordToFind[letter1]) != rune(master.WordToFind[letter2]) {
		master.UsedLetters = append(master.UsedLetters, string(rune(master.WordToFind[letter2])))
	}
	var VerifWord strings.Builder
	for _, char := range master.WordToFind {
		if master.GivenLetters[char] {
			VerifWord.WriteRune(char)
		} else {
			VerifWord.WriteString("_")
		}
	}
	master.HiddenWord = VerifWord.String()
	DisplayInterface()
}

// Permet de retourner au menu à partir de la page de jeu
func ReturnIntoMainMenu() {
	master.CurrentPage = "menu"
	master.MenuIndex = 1
	DisplayMenu()
}

// Gère les actions effectuées après avoir appuyé sur "Entrée"
func ExecIntoGame(r rune) {
	if !master.IsWriting {
		// pas assez de tentative pour un mot entier
		if master.Chances < 2 {
			color.Red("Vous n'avez plus assez de tentative pour un mot entier")
		} else {
			// on active la capture pour un mot entier
			master.IsWriting = true
			master.WritingWord = ""
			DisplayInterface()
			color.Cyan("Quel est le mot :")
			color.Yellow(master.WritingWord)
			color.Cyan("Appuyer sur entrée pour confirmer")
		}
	} else {
		// on arrête la capture de mot entier et on test le résultat
		master.IsWriting = false
		if master.WritingWord == master.WordToFind {
			// victoire
			DisplayWinScreen()
		} else {
			// raté
			master.Chances -= 2
			master.Errors += 2
			DisplayInterface()
			color.Red("Ce n'est pas le bon mot")
			if master.Chances <= 0 {
				// défaite
				DisplayLoseScreen()
			}
		}
	}
}

// Permet d'ajouter une lettre au mot qu'on écrit
func WritingWord(r rune) {
	master.WritingWord += string(rune(r))
	DisplayInterface()
	color.Cyan("Quel est le mot :")
	color.Yellow(master.WritingWord)
	color.Cyan("Appuyer sur entrée pour confirmer")
}

// Permet de tester une lettre
func TestLetter(r rune) {
	DisplayInterface()
	// on a déjà envoyé cette lettre
	if master.GivenLetters[r] {
		color.Magenta("Vous avez déjà envoyé cette lettre")
	} else if strings.ContainsRune(master.WordToFind, r) {
		// le mot contient la lettre donnée
		master.GivenLetters[r] = true
		master.UsedLetters = append(master.UsedLetters, string(rune(r)))
		var VerifWord strings.Builder
		for _, char := range master.WordToFind {
			if master.GivenLetters[char] {
				VerifWord.WriteRune(char)
			} else {
				VerifWord.WriteString("_")
			}
		}
		master.HiddenWord = VerifWord.String()
		DisplayInterface()
		if !strings.Contains(master.HiddenWord, "_") {
			// victoire
			DisplayWinScreen()
		}
	} else {
		// le mot ne contient pas la lettre donnée
		master.GivenLetters[r] = true
		master.UsedLetters = append(master.UsedLetters, string(rune(r)))
		master.Chances--
		master.Errors++
		DisplayInterface()
		if master.Chances == 0 {
			// défaite
			DisplayLoseScreen()
		}
	}
}

// Affiche les indications de victoire
func DisplayWinScreen() {
	master.HiddenWord = master.WordToFind
	DisplayInterface()
	master.CanPlay = false
	color.Green("Félicitations! Vous avez deviné le mot: %s\n", master.WordToFind)
	color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
}

// Affiche les indications de défaite
func DisplayLoseScreen() {
	DisplayInterface()
	master.CanPlay = false
	color.Red("Pendu ! Bahahaha")
	color.Yellow("Le mot était : %s", master.WordToFind)
	color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
}

// Affiche les indications de défaite
// ancienne fonction, n'est plus utilisée
func NoMoreChance() {
	master.CanPlay = false
	fmt.Println("Vous avez utilisé toute vos chances")
	color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
}

// Affiche l'interface de jeu
func DisplayInterface() {
	utils.ClearTerminal()
	data := [][]string{
		[]string{"Trouver le mot :\n" + master.HiddenWord, master.Hangman[master.Errors]},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Tentative(s) restante(s) : " + strconv.Itoa(master.Chances) + "\nLettre(s) utilisée(s)\n[ " + strings.Join(master.UsedLetters, ", ") + " ]", "Erreur(s) : " + strconv.Itoa(master.Errors)})
	table.Render()
	color.Magenta(">> Appuyer sur une touche pour la tester")
	color.Magenta(">> Appuyer sur entrée pour tester un mot")
}
