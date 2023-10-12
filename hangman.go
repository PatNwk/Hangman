package main

import (
	"fmt"
	"os"
)

func main() {
	nomfichier := "hangman.txt"
	file, err := os.ReadFile(nomfichier)
	if err != nil {
		fmt.Printf("Error reading file")
		fmt.Print("\n")
		return
	}
	fmt.Print(string(file))
}
