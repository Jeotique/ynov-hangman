package master

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
