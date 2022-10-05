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

	var Tableau = []string{"generate", "string", "number", "make", "find", "grow", "array", "node", "chicago", "list", "do", "type", "78", "before", "23", "structure", "int", "drink", "7", "set", "car", "get", "random", "are", "now", "number", "mod", "after", "file", "ynov", "print", "keep", "windows", "show", "read", "86", "hangman", "a"}

	fmt.Println(Tableau[0]) // Ecrire le premier élément du tableau

	fmt.Println(Tableau[len(Tableau)-1]) // Ecrire "l'oppose du premier" donc dernier élément du tableau

	for _, i := range Tableau {

		if i == "before" {

			fmt.Println(Tableau[14:15]) // Prendre l'élément du tableau apres le mot "before" cela désigne la string "23"

			if Tableau[23] == "are" && "are" < "now" { // à l'indice 23 du tableau nous avons "are" est il est avant le mot "maintenant"

				fmt.Println((97 + 114 + 101) / 114) // (a : 97) + (r : 114) + (e : 101) le tout diviser par (r : 114) = 2
				fmt.Println(Tableau[2])

			}
		}
	}

	return Tableau
}
