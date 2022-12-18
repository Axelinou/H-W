package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// const quote = string(byte(34))

var a []string = []string{
	" #####",
	"#    #",
	"#    #",
	"######",
	"#    #",
	"#    #",
	"      ",
}
var b []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#    #",
	"##### ",
	"      ",
}
var c []string = []string{
	" #####",
	"#     ",
	"#     ",
	"#     ",
	" #####",
	"      ",
}
var d []string = []string{
	"##### ",
	"#    #",
	"#    #",
	"#    #",
	"##### ",
	"      ",
}
var e []string = []string{
	"######",
	"#     ",
	"####  ",
	"#     ",
	"######",
	"      ",
}
var f []string = []string{
	"######",
	"#     ",
	"####  ",
	"#     ",
	"#     ",
	"      ",
}
var g []string = []string{
	" #####",
	"#     ",
	"#  ###",
	"#    #",
	" #####",
	"      ",
}
var h []string = []string{
	"#    #",
	"#    #",
	"######",
	"#    #",
	"#    #",
	"      ",
}
var i []string = []string{
	"##### ",
	"  #   ",
	"  #   ",
	"  #   ",
	"##### ",
	"      ",
}
var j []string = []string{
	"     #",
	"     #",
	"     #",
	" #   #",
	"  ##  ",
	"      ",
}
var k []string = []string{
	"#    #",
	"#  #  ",
	"##    ",
	"#  #  ",
	"#    #",
	"      ",
}
var l []string = []string{
	"#     ",
	"#     ",
	"#     ",
	"#     ",
	"######",
	"      ",
}
var m []string = []string{
	"#    #",
	"##  ##",
	"# ## #",
	"#    #",
	"#    #",
	"      ",
}
var n []string = []string{
	"#    #",
	"##   #",
	"# #  #",
	"#  # #",
	"#   ##",
	"      ",
}
var o []string = []string{
	" #### ",
	"#    #",
	"#    #",
	"#    #",
	" #### ",
	"      ",
}
var p []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#     ",
	"#     ",
	"      ",
}
var q []string = []string{
	" #### ",
	"#    #",
	"#    #",
	" #### ",
	"     # ",
	"      ",
}
var r []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#  #  ",
	"#    #",
	"      ",
}
var s []string = []string{
	" #####",
	"#     ",
	" #### ",
	"     #",
	"##### ",
	"      ",
}
var t []string = []string{
	"######",
	"  #   ",
	"  #   ",
	"  #   ",
	"  #   ",
	"      ",
}
var u []string = []string{
	"#    #",
	"#    #",
	"#    #",
	"#    #",
	" #### ",
	"      ",
}
var v []string = []string{
	"#    #",
	"#    #",
	" #  # ",
	" #  # ",
	"  ##  ",
	"      ",
}
var w []string = []string{
	"# # # ",
	"# # # ",
	"# # # ",
	"# # # ",
	" # #  ",
	"      ",
}
var x []string = []string{
	"#   # ",
	" # #  ",
	"  #   ",
	" # #  ",
	"#   # ",
	"      ",
}
var y []string = []string{
	"#    #",
	"#    #",
	" #### ",
	"  #   ",
	"  #   ",
	"      ",
}
var z []string = []string{
	"######",
	"    # ",
	"  ##  ",
	" #    ",
	"######",
	"      ",
}
var space []string = []string{
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
}
var underscore []string = []string{
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
	"######",
}

var letters [][]string = [][]string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, space, underscore} // Initialisation de l'alphabet

func PrintAscii(text string) string { // fonction permettant d'imprimer en art ASCII Art du texte en majuscule simple, un espace et un underscore
	text = strings.ToUpper(text)
	upperText := []byte{}
	for i := 0; i < len(text); i++ {
		if text[i] > 96 && text[i] < 123 {
			upperText = append(upperText, text[i]-32-65)
		} else if text[i] == 32 {
			upperText = append(upperText, 26)
		} else if text[i] == 95 {
			upperText = append(upperText, 27)
		} else {
			upperText = append(upperText, text[i]-65)
		}
	}
	asciiIndices := []int{}
	for i := 0; i < len(upperText); i++ {
		asciiIndices = append(asciiIndices, int(upperText[i]))
	}
	result := ""
	for i := 0; i < 6; i++ {
		for j := 0; j < len(asciiIndices); j++ {
			result += letters[asciiIndices[j]][i]
			result += " "
		}
		result += "\n"
	}
	return result
}

func PrintHangman(fails int) string { // permet d'afficher l'état du pendu
	data, err := ioutil.ReadFile("../text/hangman.txt")
	if err != nil {
		os.Exit(0)
	}
	hangman := string(data)
	result := "\n"
	for i := 0; i < 78; i++ {
		if fails > 10 {
			fails = 10
		}
		result += string(hangman[fails*79+i])
	}
	result += "\n"
	return result
}

func HideLetters(word, knownLetters string) string { //affiche le mot à deviner et ses indices (le mot change en fonction bonnes réponses du joueur)
	final := ""
	for i := 0; i < len(word); i++ {
		if strings.Contains(knownLetters, string(word[i])) {
			final += string(word[i])
		} else {
			final += "_"
		}
	}
	return final
}

type Hangman struct {
	guess            string
	word             string
	wrongGuesses     int
	attemptedLetters string
	Display          string
	Hangman          string
	Text             string
	difficulty       string
}

var indextmpl = template.Must(template.ParseFiles("resources/index1.html"))
var tmpl = template.Must(template.ParseFiles("resources/index.html"))
var stringData string = ""            // les données du fichier selectionné
var wordsList []string = []string{""} // les données du fichier selectionné
var word string = ""                  // mot choisi
var attemptedLetters string = ""      // lettres déja essayés
var wrongGuesses = 0                  // le nombre d'erreurs du joueur
var done bool = false                 // détermine si la partie est terminée

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	difficulty := r.FormValue("difficulty")
	if difficulty != "" {
		SetDifficulty(difficulty)

	}

	if word == "" || done {
		word = wordsList[rand.Intn(len(wordsList))]

		attemptedLetters = "" //indices
		for len(attemptedLetters) < (len(word)/2 - 1) {
			addedLetter := string(word[rand.Intn(len(word))])
			if !strings.Contains(attemptedLetters, addedLetter) {
				attemptedLetters += addedLetter
			}
		}

		wrongGuesses = 0
		done = false
	}
	guess := r.FormValue("w")

	currentText := ""
	if len(guess) > 1 {
		if word == guess {
			done = true
			attemptedLetters = guess
		} else {
			wrongGuesses += 2
			if wrongGuesses >= 11 {
				wrongGuesses = 10
			}
			currentText = fmt.Sprintf("%s n'est pas le mot... il vous reste %d essais \n", guess, (10 - wrongGuesses))
		}
	} else if len(guess) == 1 {
		if strings.Contains(attemptedLetters, guess) {
			currentText = "Vous avez déjà essayé cette lettre; Veuillez réessayer.\n\n"
		} else {
			if strings.Contains(word, guess) {
				currentText = fmt.Sprintf("%s est dans le mot! \n\n", guess)
			} else {
				currentText = fmt.Sprintf("%s n'est pas dans le mot... \n\n", guess)
				wrongGuesses += 1
			}
			attemptedLetters += guess
		}
	} else {
		currentText = fmt.Sprintf("%s n'est pas dans le mot... \n\n", r.FormValue("difficulty"))
	}
	time.Sleep(2 * time.Second)

	//verifie si le joueur à perdu ou gagné
	if HideLetters(word, attemptedLetters) == HideLetters(word, "abcdefghijklmnopqrstuvwxyz") {
		currentText = "Vous avez Gagné!" + fmt.Sprintf("Le Mot Etait %s\n", word)
		done = true
	} else if wrongGuesses >= 10 {
		currentText = "Vous avez Perdu..." + fmt.Sprintf("Le Mot Etait %s\n", word)
		done = true
	}
	fmt.Println("is there someone else ?")

	data := Hangman{
		guess:            "",
		word:             word,
		wrongGuesses:     wrongGuesses,
		attemptedLetters: attemptedLetters,
		Display:          PrintAscii(HideLetters(word, attemptedLetters)),
		Hangman:          PrintHangman(wrongGuesses),
		Text:             currentText,
		difficulty:       difficulty,
	}

	tmpl.Execute(w, data)

}

func HttpHandlerindex(w http.ResponseWriter, r *http.Request) {
	currentText := ""

	difficulty := r.FormValue("difficulty")
	if difficulty != "" {
		SetDifficulty(difficulty)
		currentText = fmt.Sprintf(" la difficulté %s est séléctionnée  \n\n", difficulty)

		http.Redirect(w, r, "/hangman", http.StatusSeeOther)

	} else {
		currentText = "Bienvenue choisissez une difficulté"
		data := Hangman{
			guess:            "",
			word:             word,
			wrongGuesses:     wrongGuesses,
			attemptedLetters: attemptedLetters,
			Display:          PrintAscii(HideLetters(word, attemptedLetters)),
			Hangman:          PrintHangman(wrongGuesses),
			Text:             currentText,
			difficulty:       difficulty,
		}
		indextmpl.Execute(w, data)

	}

}

func SetDifficulty(difficulty string) {
	data, err := ioutil.ReadFile(fmt.Sprintf("../text/%s.txt", difficulty))
	if err != nil {
		fmt.Println("Erreur : fichier introuvable")
		os.Exit(15)
	}
	stringData = string(data)
	wordsList = strings.Fields(stringData)
	done = true
}
func HttpPrint(w http.ResponseWriter, text string) {
	segmentedText := strings.Fields(text)
	for i := 0; i < len(segmentedText); i++ {
		io.WriteString(w, segmentedText[i])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dataEasy, err1 := ioutil.ReadFile("../text/easy.txt")
	if err1 != nil {
		fmt.Println("Erreur : fichier introuvable")
		os.Exit(15)
	}
	stringData = string(dataEasy)
	wordsList = strings.Fields(stringData)
	word = ""

	tmpl = template.Must(template.ParseFiles("resources/index.html"))

	styleServer := http.FileServer(http.Dir("css1"))
	http.Handle("/css1/", http.StripPrefix("/css1/", styleServer))
	fmt.Println(http.FileServer(http.Dir("css")))
	http.HandleFunc("/", HttpHandlerindex)
	http.HandleFunc("/hangman", HttpHandler)
	http.ListenAndServe(":80", nil)
}
