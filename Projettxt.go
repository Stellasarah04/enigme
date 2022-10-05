package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("Hangman.txt")

	if err != nil {
		log.Fatal(err)
	}

	ReadFile("")

	defer file.Close()
}

func ReadFile(Mot string) []string {

	var Tableau = []string{"Hangman.txt"}

	for _, i := range Tableau {
		if i == "\n" {
			Tableau = append(Tableau, Mot)
			Mot = ""
		} else {
			Mot += string(i)
		}
	}

	fmt.Println(Tableau)

	return Tableau

}

/*	for _, i := range Tableau {
		if i == "\n" {
			Tableau = append(Tableau, Mot)
			Mot = ""
		} else {
			Mot += string(i)
		}
	}

	if Mot == "before" {
		fmt.Println(string(Tableau[14]) + 1)
	}*/
