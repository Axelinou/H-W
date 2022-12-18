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

// tableau qui contient toute les lettres d'ascii art

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

func PrintAscii(text string) string { // affichage du ascii
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

func PrintHangman(fails int) string { // permet d'afficher l'état du pendu en prenant en compte le nombre d'erreurs du joueur
	data, err := ioutil.ReadFile("../text/hangman.txt") // lecture du fichier hangman comme dans le projet précèdant
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

func HideLetters(word, knownLetters string) string { //indices variables selon les réponses du joueurs
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

var defeattmpl = template.Must(template.ParseFiles("resources/defeat.html"))   //template de la page defeat
var victorytmpl = template.Must(template.ParseFiles("resources/victory.html")) //template de la page victory

var indextmpl = template.Must(template.ParseFiles("resources/index1.html")) // template de la page d'acceuil
var tmpl = template.Must(template.ParseFiles("resources/index.html"))       // template du hangman

var stringData string = ""            // les données du fichier selectionné
var wordsList []string = []string{""} // les données du fichier selectionné
var word string = ""                  // mot choisi aléatoirement
var attemptedLetters string = ""      // lettres déja essayés
var wrongGuesses = 0                  // le nombre d'erreurs du joueur
var done bool = false                 // détermine si la partie est terminée

func HttpHandler(w http.ResponseWriter, r *http.Request) { // fonction principale du hangman

	difficulty := r.FormValue("difficulty") // recupère la valeur des boutons difficulté
	if difficulty != "" {
		SetDifficulty(difficulty) // lecture du fichier text de meme nom

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
	guess := r.FormValue("w") // valeur entrée par l'utilisateur dans le formulaire

	currentText := ""
	if len(guess) > 1 {
		if word == guess { //si le mot correspond à au mot entré par l'utilisateur
			done = true
			attemptedLetters = guess
		} else { // si c'est pas le cas on ajoute 2 à ses erreurs
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
		currentText = fmt.Sprintf(" Il vous reste %d tentatives bonne chance ! \n\n", (10 - wrongGuesses)) //message par défaut
	}
	time.Sleep(2 * time.Second)

	//verifie si le joueur à perdu ou gagné
	if HideLetters(word, attemptedLetters) == HideLetters(word, "abcdefghijklmnopqrstuvwxyz") {
		currentText = "Vous avez Gagné!" + fmt.Sprintf("Le Mot Etait %s\n", word)
		done = true
		http.Redirect(w, r, "/victory", http.StatusSeeOther) //redirection vers la page de victoire

	} else if wrongGuesses >= 10 {
		currentText = "Vous avez Perdu..." + fmt.Sprintf("Le Mot Etait %s\n", word)
		done = true
		http.Redirect(w, r, "/defeat", http.StatusSeeOther) //redirection vers la page de défaite
	}

	data := Hangman{ //valeur de la template
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

func HttpHandlerindex(w http.ResponseWriter, r *http.Request) { // fonction principale de la page d'acceuil
	currentText := ""

	difficulty := r.FormValue("difficulty")
	if difficulty != "" {
		SetDifficulty(difficulty)

		http.Redirect(w, r, "/hangman", http.StatusSeeOther) // redirection vers la page du hangman

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

func DefeatHttpHandler(w http.ResponseWriter, r *http.Request) {

	data := Hangman{
		guess:            "",
		word:             word,
		wrongGuesses:     wrongGuesses,
		attemptedLetters: attemptedLetters,
		Display:          PrintAscii(HideLetters(word, attemptedLetters)),
		Hangman:          PrintHangman(wrongGuesses),
	}

	defeattmpl.Execute(w, data)
}
func VictoryHttpHandler(w http.ResponseWriter, r *http.Request) {

	data := Hangman{
		guess:            "",
		word:             word,
		wrongGuesses:     wrongGuesses,
		attemptedLetters: attemptedLetters,
		Display:          PrintAscii(HideLetters(word, attemptedLetters)),
		Hangman:          PrintHangman(wrongGuesses),
	}

	victorytmpl.Execute(w, data)
}
func SetDifficulty(difficulty string) { //fonction permettant  de selectionner le fichier avec la difficulté choisie
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

	styleServer := http.FileServer(http.Dir("css1")) //permet d'indiquer au serveur ou trouver le css et les images
	http.Handle("/css1/", http.StripPrefix("/css1/", styleServer))

	http.HandleFunc("/", HttpHandlerindex)          //adresse url de la page d'acceuil
	http.HandleFunc("/hangman", HttpHandler)        //adresse url du hangman
	http.HandleFunc("/defeat", DefeatHttpHandler)   //adresse url de la page de defaite
	http.HandleFunc("/victory", VictoryHttpHandler) //adresse url de la page de victoire
	http.ListenAndServe(":80", nil)                 //utilisation du port 80
}
