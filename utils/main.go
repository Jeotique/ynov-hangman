package utils

import (
	"math/rand"
	"os"
	"os/exec"
	"projet/master"
	"runtime"
	"strings"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Supprime le contenu de la console
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

// Obient un mot aléatoire parmis le txt donné
func GetRandomWord() string {
	lines, _ := os.ReadFile(master.TextBook)
	all := strings.Split(string(lines), "\n")
	word := all[rand.Intn(len(all))]
	var finalWord string
	for _, i := range word {
		if i != 13 {
			finalWord += string(rune(i))
		}
	}
	finalWord = strings.ReplaceAll(finalWord, " ", "")
	finalWord = strings.ToLower(finalWord)
	return strings.Trim(finalWord, "\n")
}
