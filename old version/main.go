package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mattn/go-tty"
	"github.com/olekukonko/tablewriter"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// toute les variables du programme
var (
	WordToFind   string
	HiddenWord   string
	GivenLetters = make(map[rune]bool)
	UsedLetters  []string
	Chances      = 10
	Errors       = 0
	CanPlay      = false
	TextBook     = "mots.txt"
	IsWriting    = false
	WritingWord  string
	CurrentPage  = "menu"
	MenuIndex    = 1
)

// dessin ASCII ART du pendu
var Hangman = []string{
	"\n\n\n\n\n\n",
	"\n\n\n\n\n\n=========",
	"      |  \n      |  \n      |  \n      |  \n      |  \n=========",
	"  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========",
	"  +---+\n  |   |\n      |\n      |\n      |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n      |\n      |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n  |   |\n      |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n /|   |\n      |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n /|\\  |\n      |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n /|\\  |\n /    |\n      |\n=========",
	"  +---+\n  |   |\n  O   |\n /|\\  |\n / \\  |\n      |\n=========",
}

func main() {
	// change la cible du .txt
	if len(os.Args) >= 2 && os.Args[1] != "" {
		TextBook = os.Args[1]
	}
	//NewGame()
	DisplayMenu()
	StartListening()
}

func DisplayMenu() {
	ClearTerminal()
	switch MenuIndex {
	case 1:
		data := [][]string{
			[]string{"> Jouer"}, []string{"Changer de dictionnaire | " + TextBook}, []string{"Quitter"},
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
			[]string{"Jouer"}, []string{"> Changer de dictionnaire | " + TextBook}, []string{"Quitter"},
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
			[]string{"Jouer"}, []string{"Changer de dictionnaire | " + TextBook}, []string{"> Quitter"},
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

func SelectDictionnary() {
	ClearTerminal()
	color.Cyan("Quel est le nom du fichier :")
	color.Yellow(WritingWord)
	color.Cyan("Appuyer sur entrée pour confirmer")
}

// Initialise une nouvelle partie
func NewGame() {
	CanPlay = true
	Chances = 10
	Errors = 0
	WordToFind = GetRandomWord()
	WordToFind = strings.ReplaceAll(WordToFind, "\n", "")
	HiddenWord = strings.Repeat("_", len(WordToFind))
	GivenLetters = make(map[rune]bool)
	UsedLetters = []string{}
	letter1 := rand.Intn(len(WordToFind))
	letter2 := rand.Intn(len(WordToFind))
	GivenLetters[rune(WordToFind[letter1])] = true
	GivenLetters[rune(WordToFind[letter2])] = true
	UsedLetters = append(UsedLetters, string(rune(WordToFind[letter1])))
	if rune(WordToFind[letter1]) != rune(WordToFind[letter2]) {
		UsedLetters = append(UsedLetters, string(rune(WordToFind[letter2])))
	}
	var VerifWord strings.Builder
	for _, char := range WordToFind {
		if GivenLetters[char] {
			VerifWord.WriteRune(char)
		} else {
			VerifWord.WriteString("_")
		}
	}
	HiddenWord = VerifWord.String()
	DisplayInterface()
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
		if !CanPlay && CurrentPage == "menu" {
			if r == 65 {
				if MenuIndex == 1 {
					MenuIndex = 3
				} else {
					MenuIndex--
				}
				DisplayMenu()
			} else if r == 66 {
				if MenuIndex == 3 {
					MenuIndex = 1
				} else {
					MenuIndex++
				}
				DisplayMenu()
			}
			if r == 32 {
				switch MenuIndex {
				case 1:
					MenuIndex = 0
					CurrentPage = "ingame"
					NewGame()
					break
				case 2:
					IsWriting = true
					WritingWord = ""
					CurrentPage = "select"
					SelectDictionnary()
					break
				case 3:
					os.Exit(0)
				}
			}
		} else if !CanPlay && CurrentPage == "select" {
			if r == 13 { // touche entrée
				IsWriting = false
				TextBook = WritingWord
				WritingWord = ""
				CurrentPage = "menu"
				DisplayMenu()
			} else if r >= 97 && r <= 122 || r == 46 || r == 95 {
				WritingWord += string(rune(r))
				SelectDictionnary()
			}
		}

		if r == 8 { // touche effacer
			// permet de relancer une partie si la précédente est terminée
			if !CanPlay && CurrentPage == "ingame" {
				CurrentPage = "menu"
				MenuIndex = 1
				DisplayMenu()
			}
		}
		if r >= 97 && r <= 122 || r == 45 || r == 13 {
			// si les touches sont des lettres, tiret, entrée
			if CanPlay {
				// la partie est en cours, on capture donc les inputs
				if r == 13 { // touche entrée
					if !IsWriting {
						// pas assez de tentative pour un mot entier
						if Chances < 2 {
							color.Red("Vous n'avez plus assez de tentative pour un mot entier")
						} else {
							// on active la capture pour un mot entier
							IsWriting = true
							WritingWord = ""
							DisplayInterface()
							color.Cyan("Quel est le mot :")
							color.Yellow(WritingWord)
							color.Cyan("Appuyer sur entrée pour confirmer")
						}
					} else {
						// on arrête la capture de mot entier et on test le résultat
						IsWriting = false
						if WritingWord == WordToFind {
							// victoire
							HiddenWord = WordToFind
							DisplayInterface()
							CanPlay = false
							color.Green("Félicitations! Vous avez deviné le mot: %s\n", WordToFind)
							color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
						} else {
							// raté
							Chances -= 2
							Errors += 2
							DisplayInterface()
							color.Red("Ce n'est pas le bon mot")
							if Chances <= 0 {
								// défaite
								DisplayInterface()
								CanPlay = false
								color.Red("Pendu ! Bahahaha")
								color.Yellow("Le mot était : %s", WordToFind)
								color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
							}
						}
					}
				} else {
					if IsWriting {
						// on capture les touches pour un mot entier donc on ajoute la lettre
						WritingWord += string(rune(r))
						DisplayInterface()
						color.Cyan("Quel est le mot :")
						color.Yellow(WritingWord)
						color.Cyan("Appuyer sur entrée pour confirmer")
					} else {
						// on test la lettre si il nous reste des tentatives
						if Chances > 0 {
							DisplayInterface()
							// on a déjà envoyé cette lettre
							if GivenLetters[r] {
								color.Magenta("Vous avez déjà envoyé cette lettre")
							} else if strings.ContainsRune(WordToFind, r) {
								// le mot contient la lettre donnée
								GivenLetters[r] = true
								UsedLetters = append(UsedLetters, string(rune(r)))
								var VerifWord strings.Builder
								for _, char := range WordToFind {
									if GivenLetters[char] {
										VerifWord.WriteRune(char)
									} else {
										VerifWord.WriteString("_")
									}
								}
								HiddenWord = VerifWord.String()
								DisplayInterface()
								if !strings.Contains(HiddenWord, "_") {
									// victoire
									CanPlay = false
									color.Green("Félicitations! Vous avez deviné le mot: %s\n", WordToFind)
									color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
								}
							} else {
								// le mot ne contient pas la lettre donnée
								GivenLetters[r] = true
								UsedLetters = append(UsedLetters, string(rune(r)))
								Chances--
								Errors++
								DisplayInterface()
								if Chances == 0 {
									// défaite
									CanPlay = false
									color.Red("Pendu ! Bahahaha")
									color.Yellow("Le mot était : %s", WordToFind)
									color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
								}
							}
						} else {
							CanPlay = false
							fmt.Println("Vous avez utilisé toute vos chances")
							color.Cyan("Appuyer sur la touche effacer pour revenir au menu")
						}
					}
				}
			}
		}
	}
}

// Affiche l'interface de jeu
func DisplayInterface() {
	ClearTerminal()
	data := [][]string{
		[]string{"Trouver le mot :\n" + HiddenWord, Hangman[Errors]},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Tentative(s) restante(s) : " + strconv.Itoa(Chances) + "\nLettre(s) utilisée(s)\n[ " + strings.Join(UsedLetters, ", ") + " ]", "Erreur(s) : " + strconv.Itoa(Errors)})
	table.Render()
	color.Magenta(">> Appuyer sur une touche pour la tester")
	color.Magenta(">> Appuyer sur entrée pour tester un mot")
}

// Obient un mot aléatoire parmis le txt donné
func GetRandomWord() string {
	lines, _ := os.ReadFile(TextBook)
	all := strings.Split(string(lines), "\n")
	word := all[rand.Intn(len(all))]
	var finalWord string
	for _, i := range word {
		if i != 13 {
			finalWord += string(rune(i))
		}
	}
	return strings.Trim(finalWord, "\n")
}

// UTILS

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
