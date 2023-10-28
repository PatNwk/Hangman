package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func lan() string {
	scan1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Now choose the language:")
	fmt.Println()
	fmt.Println("For English type 1")
	fmt.Println("For French type 2")
	scan1.Scan()
	rep1 := scan1.Text()
	return rep1

}

// Création du pendu
func pendue(nomFichier string) []string {
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil
	}
	defer file.Close()
	var image []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		image = append(image, scanner.Text())
	}
	pendu := strings.Split(strings.Join(image, "\n"), "=========")

	return pendu
}

func main() {
	var listeeeee []string
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("To start the game type 'Start'")
	scan.Scan()
	rep := scan.Text()
	if rep == "Start" {
		var nomfichierr string
		oui := lan()
		if oui == "1" {
			nomfichierr = "words2.txt"
		} else if oui == "2" {
			nomfichierr = "words.txt"
		}
		file, err := os.ReadFile(nomfichierr)
		if err != nil {
			fmt.Printf("file name missing")
			fmt.Print("\n")
			return
		}
		pendu := pendue("hangman.txt")
		text := string(file)
		words := strings.Fields(text)
		for i := 0; i < len(words); i++ {
			words[i] = strings.Trim(words[i], `",`)
		}

		rand.Seed(time.Now().Unix())
		randomIndex := rand.Intn(len(words))
		randomWord := words[randomIndex]
		//fmt.Println(randomWord)
		ind := len(randomWord)/2 - 1
		lettre := string(randomWord[ind])
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		if oui == "1" {
			fmt.Println("The game starts !")
		} else if oui == "2" {
			fmt.Println("Le jeu commence !")
		}
		// Affichez le mot partiellement masqué
		fmt.Println()
		motMasque := make([]string, len(randomWord))
		for i := 0; i < len(randomWord); i++ {
			if i == ind {
				motMasque[i] = lettre
			} else {
				motMasque[i] = "_"
			}
		}
		fmt.Println(strings.Join(motMasque, " "))
		lettreTrouvee := false
		scanner := bufio.NewScanner(os.Stdin)
		// Compteur de vie
		compt := len(pendu) - 1
		ratio := 0

		for {
			fmt.Println()
			fmt.Println("--------------------------------------------------")
			if oui == "1" {
				fmt.Print("Choose: ")
			} else if oui == "2" {
				fmt.Print("Lettre choisie: ")
			}
			scanner.Scan()
			reponse := scanner.Text()
			reponse = strings.ToLower(reponse)
			fmt.Println()
			listeeeee = append(listeeeee, reponse)
			if len(reponse) > 1 {
				if oui == "1" {
					fmt.Println("Please type only one letter")
				} else {
					fmt.Println("Veuillez taper qu'une seule lettre")
				}
			}
			if oui == "1" {
				fmt.Println("Letters already typed: ", listeeeee)
			} else if oui == "2" {
				fmt.Println("Lettres déjà tapées: ", listeeeee)
			}
			// Vérifiez si la lettre est présente dans le mot
			lettreTrouvee = false
			res := false
			for i := 0; i < len(randomWord); i++ {
				if reponse == string(randomWord[i]) && len(reponse) < 2 {
					motMasque[i] = reponse
					lettreTrouvee = true
					res = true
				}
			}
			// Si la lettre n'est pas dans le mot
			if !lettreTrouvee && len(reponse) < 2 {
				compt = compt - 1
				if oui == "1" {
					fmt.Println("Not present in the word", compt, "attempts remaining")
				} else if oui == "2" {
					fmt.Println("Lettre non présente dans le mot", compt, "tentatives restantes")
				}
			}
			// Affichage du pendu
			if compt >= 0 && compt < len(pendu)-1 && len(reponse) < 2 {
				if !res {
					pendu[ratio] += "========="
					fmt.Println(pendu[ratio])
					ratio = ratio + 1
				}
			}
			// Affichez le mot mis à jour
			fmt.Println(strings.Join(motMasque, " "))
			fmt.Println()
			if !strings.Contains(strings.Join(motMasque, ""), "_") {
				if oui == "1" {
					fmt.Println("Congrats you find the word !")
				} else if oui == "2" {
					fmt.Println("Félicitation vous avez trouvé le mot !")
				}
				break
			}
			// Réponse si on a perdu
			if compt == 0 {
				fmt.Println()
				if oui == "1" {
					fmt.Println("You loose !")
					fmt.Println("The word was", randomWord, "!")
					fmt.Println(pendu[9])
				} else if oui == "2" {
					fmt.Println("vous avez perdu !")
					fmt.Println("Le mot était", randomWord, "!")
					fmt.Println(pendu[9])
				}
				break
			}
		}
		// Demande si il veut rejouer
		if oui == "2" {
			fmt.Println("Voulez-vous rejouer ?")
			fmt.Println("Tapez '1' pour rejouer, '2' pour quitter.")
		} else {
			fmt.Println("Do you want to play again ?")
			fmt.Println("Type '1' to play again, '2' to exit.")
		}
		scanner.Scan()
		rejouer := scanner.Text()

		if rejouer == "1" {
			main()
		} else {
			if oui == "2" {
				fmt.Println("Merci d'avoir joué !")
			} else {
				fmt.Println("Thanks for playing !")
			}
		}
	} else {
		fmt.Println("You typed the word wrong")
		main()
	}
}
