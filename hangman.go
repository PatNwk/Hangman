package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Création du pendu
func pendue(nomFichier string) []string {
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil
	}
	defer file.Close()
	var positions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		positions = append(positions, scanner.Text())
	}
	pendu := strings.Split(strings.Join(positions, "\n"), "=========")

	return pendu
}
func main() {
	nomfichier := os.Args[1]
	file, err := os.ReadFile(nomfichier)
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
		fmt.Print("Choose: ")
		scanner.Scan()
		reponse := scanner.Text()

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
			fmt.Println("Not present in the word", compt, "attempts remaining")
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
			break
		}
		// Réponse si on a perdu
		if compt == 0 {
			fmt.Println("You loose !")
			fmt.Println("The word was", randomWord, "!")
			fmt.Println(pendu[9])
			break
		}
	}
}
