package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	"unicode"
)

type Hangman struct {
	lentab          int    //longueur du tableau
	lenword         int    // variable qui stocke la longueur du mot choisi
	userletter      string // variable qui stocke la lettre entrée par l'utilisateur
	testsctword     string // variable qui stocke le mot choisi
	nbinwordpresent int    // variable qui stocke le nombre de fois que la lettre donné par l'utilisateur est présente dans le mot choisi
	slctletter      string // variable qui stocke la lettre choisie aléatoirement pour etre affichée
	userletter1     string //
	userletter2     string //
	userletter3     string //
	userletter4     string //
	userletter5     string //
	userletter6     string //
	userletter7     string //
	userletter8     string //
	userletter9     string //
	userletter10    string //
	userletter11    string //
	userletter12    string
	userletter13    string
	userletter14    string
	userletter15    string // variable qui stocke la lettre entrée par l'utilisateur
	userletter16    string
	turn            int // variable qui stocke le nombre de tour
	riddletab       []string
	usedletters     []string
	indexuserletter int
	iswrong         bool // variable qui stocke si l'utilisateur a commis une faute
	attempt         int
	hasused         bool
	isanusedletter  int
}

func playgame() {
	var hangman1 Hangman
	hangman1.lentab = 0
	hangman1.lenword = 0
	hangman1.attempt = 10
	fmt.Println("Les mots sont en français sauf un seul, vous avez 10 tentatives Bonne chance !")
	hangman("word.txt", Hangman{lentab: 0, lenword: 0, attempt: 10}) // Appelle la fonction qui lit les mot du fichier texte "word"
	//			 ⬆️
	// veuillez remplacer ce chemin par celui de votre fichier word.txt (à noter que cette manipulation n'est pas nécessaire si vous exécutez ce fichier dans le cmd)

}

func IsUpper(lettre string) bool {

	for _, letter := range lettre { // Boucle qui vérifie si la lettre entrée est une majuscule
		if unicode.IsUpper(letter) {
			return true
		} else {
			fmt.Println("Votre lettre doit etre en majuscule")
			return false
		}
	}

	return false
}

func InputUser(hangman1 Hangman) {
	if hangman1.iswrong == true {
		fmt.Println("                    ")
		fmt.Println(hangman1.riddletab) // permet de ne pas surcharger d'infoermations  l'utilisateur
	}
	var Mo string // variable qui stocke la variable entrée par l'utilisateur
	fmt.Println("                                   ")
	fmt.Println("Choissisez une lettre (ou un mot):")
	fmt.Scanln(&Mo)
	hangman1.userletter = Mo

	if len(Mo) == 1 { // si la lettre est bien unique

		if IsUpper(Mo) == true {
			fmt.Println("Vous avez entré la lettre :", Mo, "")
			hangman1.usedletters = append(hangman1.usedletters, Mo)
			for i := 0; i < len(hangman1.usedletters); i++ {
				if Mo == string(hangman1.usedletters[i]) {

					hangman1.isanusedletter = hangman1.isanusedletter + 1 // on vérifie les occurences dans le tableau des lettres utilisées par l'utilisateur

				}
				if hangman1.isanusedletter > 1 { //si il a plus d'une occurence
					fmt.Println("Vous avez déjà été utilisé cette lettre")
					hangman1.usedletters = append(hangman1.usedletters[:i], hangman1.usedletters[i+1:]...) // on supprime la lettre du tableau des lettres utilisées
					hangman1.isanusedletter = 0                                                            //remise à zéro
					InputUser(hangman1)

				}

			}
			hangman1.isanusedletter = 0
			lettercheck(hangman1)

		} else {
			InputUser(hangman1)
		}
	} else {

		if len(Mo) == 0 {
			fmt.Println("❌Caractère non valide❌")
			InputUser(hangman1)

		}

		if len(Mo) > len(hangman1.testsctword) {
			fmt.Println("❌Votre mot est plus long", "(", len(Mo), " lettres)", " que le mot à trouver", "(", len(hangman1.testsctword), " lettres)", "❌")
			InputUser(hangman1)
		}

		for i := range hangman1.userletter { // parcours du mot choisi

			if string(hangman1.testsctword[i]) == string(hangman1.userletter[i]) { // si la lettre i du mot choisi est égale à la lettre i du mot entré par l'utilisateur

				hangman1.nbinwordpresent = hangman1.nbinwordpresent + 1 //on ajoute 1 au compteur de présence de lettre
			}

			if hangman1.nbinwordpresent == len(hangman1.testsctword)-1 { // si le compteur de présence de lettre est égal à la longueur du mot choisi
				fmt.Println("Vous avez gagné ! Bravo !")
				os.Exit(0) // on quitte le programme
			}

		}
		fmt.Println("❌Ce n'est pas le mot recherché❌")

		hangman1.attempt = hangman1.attempt - 2
		fmt.Println("Il vous reste", hangman1.attempt, "tentatives")
		if hangman1.attempt <= 0 {
			fmt.Println("Vous avez perdu :  José est mort")
			os.Exit(0)
		}
		hangman1.nbinwordpresent = 0 // remise à zéro du compteur de présence de lettre
		InputUser(hangman1)
	}
}

func hangman(file string, hangman1 Hangman) {
	data, err := ioutil.ReadFile(file) // lire le fichier text.txt
	if err != nil {                    // si il y a une erreur

		fmt.Println("_______________________________________________")
		fmt.Println("                                              ")
		fmt.Println("                                              ")
		fmt.Println("⚠️Erreur inattendue:  fermeture du programme⚠️")
		fmt.Println("                                              ")
		fmt.Println("_______________________________________________")

		fmt.Println("Cause:", err) //afficher la cause de l'erreur/panic
		fmt.Println()
		os.Exit(1) // fermeture du code  avec  le code de sortie  1 (le code de sortie 0 signifie la bon execution du programme), il est aussi possible d'utiliser panic() à la place

	}

	//fmt.Println(string(data))
	//fmt.Println(len(data))
	var arr []string
	var tmp string

	for _, _byte := range data { // lecture des caractères du fichier texte

		if _byte == '\n' { //si byte est un saut de ligne /espace
			arr = append(arr, tmp) //on l'ajoute au  tableau arr
			tmp = ""
		} else { //sinon on concatene
			tmp += string(_byte)
		}
	}
	rand.Seed(time.Now().Unix()) //etablissement d'une "graine" aléatoire
	hangman1.lentab = len(arr)   //longueur du tableau

	var randword = rand.Intn(hangman1.lentab) // mot aléatoire choisi parmi les mots du tableau arr

	var selctword = arr[randword]    // le mot selectionné
	hangman1.testsctword = selctword // assignation au champ dédié
	var lenselctword = len(arr[randword])

	for i := 0; i < len(selctword)-1; i++ {
		hangman1.lenword = i

	}

	hangman1.lenword = hangman1.lenword + lenselctword
	hangman1.lenword = hangman1.lenword / 2
	var randletter = rand.Intn(hangman1.lenword)
	hangman1.slctletter = string(selctword[randletter])

	var mysterytab []string
	for i := 0; i < len(selctword)-1; i++ { // boucle qui genere le tableau contenant la lettre indice et les lettre à trouver cachées par les enderscore
		if i == randletter {
			mysterytab = append(mysterytab, string(selctword[randletter])) // ajout  de la lettre indice au tablea
		} else {
			mysterytab = append(mysterytab, "_") // ajout des enderscore qui cachent les lettres à trouver
		}
	}

	fmt.Println(mysterytab) // le tableau
	hangman1.riddletab = mysterytab
	InputUser(hangman1) //appel de la fonction qui s'occupe de l'entrée utilisateur
}
func lettercheck(hangman1 Hangman) {

	var resultab []string
	for i := range hangman1.testsctword { // parcours du mot choisi

		if string(hangman1.testsctword[i]) == hangman1.userletter || string(hangman1.testsctword[i]) == hangman1.slctletter || string(hangman1.testsctword[i]) == hangman1.userletter1 || string(hangman1.testsctword[i]) == hangman1.userletter2 || string(hangman1.testsctword[i]) == hangman1.userletter3 || string(hangman1.testsctword[i]) == hangman1.userletter4 || string(hangman1.testsctword[i]) == hangman1.userletter5 || string(hangman1.testsctword[i]) == hangman1.userletter6 || string(hangman1.testsctword[i]) == hangman1.userletter7 || string(hangman1.testsctword[i]) == hangman1.userletter8 || string(hangman1.testsctword[i]) == hangman1.userletter9 || string(hangman1.testsctword[i]) == hangman1.userletter10 {
			resultab = append(resultab, string(hangman1.testsctword[i])) // si est l'une des lettre entrée par l'utilisateur et stockées dans les diffrentes variable  on le rajoute au tableau contenant la lettre indice et les lettre à trouver cachées par les enderscore

			if string(hangman1.testsctword[i]) == hangman1.userletter { // si  i est la lettre indice on montre toutes ses ocurences
				hangman1.nbinwordpresent = hangman1.nbinwordpresent + 1
			}
			hangman1.indexuserletter = i

		} else { // sinon on ajoute des enerscores des lettre non trouvées
			resultab = append(resultab, "_")
		}
		if len(resultab) == len(hangman1.testsctword) { // permet de résoudre un problème d'affichage qui affichait un enderscore en trop
			resultab = append(resultab[:i], resultab[i+1:]...)
		}
	}
	var iscorrect int //	 variable qui compte les lettres trouvées dans le mot
	for g := range resultab {
		if resultab[g] != "_" { // si g est différent d'une enderscore, si c'est une lettre

			iscorrect = iscorrect + 1 // Une autre lettre a été trouvée

		}

	}
	if iscorrect >= len(resultab) { // si toute les lettre ont été trouvées // si le nombre de lettre trouvées = à celui de la longueur de resultab
		fmt.Println(resultab)
		win(hangman1) //appel de la fonction de victoire
	}

	if hangman1.nbinwordpresent >= 1 { // si il n'y a qu'une seule lettre trouvée  dans le mot
		if hangman1.userletter == hangman1.slctletter { //si la lettre entré par l'utilisateur est la lettre indice
			if hangman1.nbinwordpresent == 1 {
				fmt.Println("❌La lettre est déja trouvée car donnée au début❌")
				hangman1.nbinwordpresent = 0
				hangman1.iswrong = true
				hangman1.attempt = hangman1.attempt - 1
				PositionHangman("hangman.txt", Hangman{attempt: hangman1.attempt, lentab: 0, lenword: 0})

				fmt.Println("Il vous reste", hangman1.attempt, "tentatives")
				fmt.Println("Vous avez déjà utilisé les lettres suivantes:", hangman1.usedletters)
				InputUser(hangman1)
			}
			if hangman1.nbinwordpresent > 1 && hangman1.hasused == false { // si la lettre donnée est présente à plusieur occurences et la deuxième condition empeche des abus

				fmt.Println(resultab)
				hangman1.nbinwordpresent = 0
				hangman1.iswrong = false // l'utilisateur n'a techniquement pas faux
				hangman1.hasused = true  // empeche des abus
				InputUser(hangman1)
			}
		} else { // si il y
			fmt.Println("✅La lettre est bien présente✅")
			hangman1.iswrong = false          // l'utilisateur a trouvé une lettre (permet d'eviter un bug dans laquel n'impore entrée de l'utilisateur est fausse)
			hangman1.turn = hangman1.turn + 1 // ajoute un tour qui permet de mettre en mémoire un lettre à chaque tour
			switch hangman1.turn {            // met en mémoire un mot en fonction de son tour
			case 1:
				{
					hangman1.userletter1 = hangman1.userletter
				}
			case 2:
				{
					hangman1.userletter2 = hangman1.userletter
				}
			case 3:
				{
					hangman1.userletter3 = hangman1.userletter
				}
			case 4:
				{
					hangman1.userletter4 = hangman1.userletter
				}
			case 5:
				{
					hangman1.userletter5 = hangman1.userletter
				}
			case 6:
				{
					hangman1.userletter6 = hangman1.userletter
				}
			case 7:
				{
					hangman1.userletter7 = hangman1.userletter
				}
			case 8:
				{
					hangman1.userletter8 = hangman1.userletter
				}
			case 9:
				{
					hangman1.userletter9 = hangman1.userletter
				}
			case 10:
				{
					hangman1.userletter10 = hangman1.userletter
				}
			case 11:
				{
					hangman1.userletter11 = hangman1.userletter
				}
			case 12:
				{
					hangman1.userletter12 = hangman1.userletter
				}
			case 13:
				{
					hangman1.userletter13 = hangman1.userletter
				}
			case 14:
				{
					hangman1.userletter14 = hangman1.userletter
				}
			case 15:
				{
					hangman1.userletter15 = hangman1.userletter
				}
			case 16:
				{
					hangman1.userletter16 = hangman1.userletter
				}
			}

			hangman1.nbinwordpresent = 0 //réinitialise le compteur pour eviter des comptages abusifs
			fmt.Print(resultab)
			InputUser(hangman1)

		}
	} else {
		fmt.Println("❌La lettre n'est pas dans le mot❌")
		hangman1.attempt = hangman1.attempt - 1 // perte d'une tentative
		if hangman1.attempt > 0 {               // si le joueur a au moins 1 tentative
			fmt.Println("Il vous reste", hangman1.attempt, "tentatives")
			fmt.Println("Vous avez déjà utilisé les lettres suivantes:", hangman1.usedletters)
		}
		hangman1.iswrong = true                                                                   // le joueur s'est trompé
		PositionHangman("hangman.txt", Hangman{attempt: hangman1.attempt, lentab: 0, lenword: 0}) //lecture du fichier contenant les positions du pendu
		hangman1.riddletab = resultab                                                             //permet d'eviter les limite d'une variable locale
		InputUser(hangman1)                                                                       // appel de la fonction qui gere les entrées de l'utilisateur
	}

}
func PositionHangman(hangman string, hangman1 Hangman) { //fonction qui lie le fichier text des positions du hangman

	data, err := ioutil.ReadFile(hangman) // lire le fichier Position.txt
	if err != nil {                       // en cas d'erreur de liecture du fichier
		//fmt.Println(fmt.Println("there is an error:", err))
		fmt.Println("_______________________________________________")
		fmt.Println("                                              ")
		fmt.Println("                                              ")
		fmt.Println("⚠️Erreur inattendue:  fermeture du programme⚠️")
		fmt.Println("                                              ")
		fmt.Println("_______________________________________________")

		fmt.Println("Cause:", err)
		os.Exit(1)
	}

	var arr []string
	var tmp string

	for _, _byte := range data { // boucle qui permet d'ajouter ce qu'il y a dans le fichier texte dans le tableau

		if _byte == '.' { // si le caractère est un .
			arr = append(arr, tmp) // on l'ajoute au tableau arr
			tmp = ""
		} else {
			tmp += string(_byte) // sinon on concatene
		}
	}

	fmt.Println(arr[hangman1.attempt]) // si le joueur n'a plus de tentatives
	if hangman1.attempt == 0 {
		lose(hangman1) // appel de la fonction défaite

	}
}

func win(hangman1 Hangman) {

	fmt.Println("VOUS AVEZ GAGNÉ  félicitations !")
	fmt.Println("le mot était", hangman1.testsctword)
	os.Exit(0) // sortie du programme le code de sortie zéro indiquant une bonne éxécution

}

func lose(hangman1 Hangman) {
	fmt.Println("VOUS AVEZ PERDU : José est mort !")
	os.Exit(0) // sortie du programme le code de sortie zéro indiquant une bonne éxécution

}
