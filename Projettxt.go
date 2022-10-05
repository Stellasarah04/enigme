package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, _ := os.Open("Hangman.txt")

	data, err := ioutil.ReadFile("Hangman.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	defer file.Close()
}

func ReadFile(Mot string) []string {

	var Tableau []string

	fmt.Println(Tableau[0])
	fmt.Println(Tableau[len(Tableau)-1])

	for _, i := range Tableau {
		if i == "\n" {
			Tableau = append(Tableau, Mot)
			Mot = ""
		}
	}
	return Tableau

}
