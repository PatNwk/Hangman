package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func langue() string {
	scan1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Now choose the language:")
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
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("To start the game type Start")
	scan.Scan()
	rep := scan.Text()
	if rep == "Start" {
		var nomfichierr string
		oui := langue()
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

		// Affichez le mot partiellement masqué
		fmt.Println()
		if oui == "1" {
			fmt.Println("The game starts !")
		} else if oui == "2" {
			fmt.Println("Le jeu commence !")
		}
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
			fmt.Println()
			// Vérifiez si la lettre est présente dans le mot
			lettreTrouvee = false
			res := false
			for i := 0; i < len(randomWord); i++ {
				if reponse == string(randomWord[i]) {
					motMasque[i] = reponse
					lettreTrouvee = true
					res = true
				}
			}
			// Si la lettre n'est pas dans le mot
			if !lettreTrouvee {
				compt = compt - 1
				if oui == "1" {
					fmt.Println("Not present in the word", compt, "attempts remaining")
				} else if oui == "2" {
					fmt.Println("Lettre non présente dans le mot", compt, "tentatives restantes")
				}
			}
			// Affichage du pendu
			if compt >= 0 && compt < len(pendu)-1 {
				if !res {
					pendu[ratio] += "========="
					fmt.Println(pendu[ratio])
					ratio = ratio + 1
				}
			}

			// Affichez le mot mis à jour
			fmt.Println(strings.Join(motMasque, " "))

			if !strings.Contains(strings.Join(motMasque, ""), "_") {
				fmt.Println("Congrats !")
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
	} else {
		fmt.Println("You typed the word wrong")
		main()
	}
}
