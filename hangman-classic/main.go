package main

import (
	"fmt"
	hangman "hangman/hangman"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func helper() {
	fmt.Println("\nUsage: ./<program_name> [word_list_file]")
	fmt.Println("The goal of the game is to find the word")
	fmt.Println("choose letter and find the word")
	fmt.Println("Good luck !!!")
	fmt.Println("")
}

func main() {
	a := 10
	tmp := os.Args[0:]
	win := 0
	var word, hiddenword string
	if len(tmp) != 2 && (os.Args[1] != "--startWith" && os.Args[2] != "save.txt") {
		fmt.Println("Argument non valide")
		return
	} else if os.Args[1] == "-h" {
		helper()
		return
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
			var err error
			data := hangman.MakeWordTab("save.txt")
			word = data[1]
			hiddenword = data[0]
			a, err = strconv.Atoi(data[2])
			if err != nil {
				return
			}
			for x := 0; x < len(hiddenword); x++ {
				fmt.Print(string(hiddenword[x]))
				fmt.Print(" ")
			}
		}
	} else {
		path := tmp[1]
		tab := hangman.MakeWordTab(path)
		rand.Seed(time.Now().UnixNano())
		rand := rand.Intn(len(tab) - 1)
		word = tab[rand]
		fmt.Println("")
		hiddenword = string(hangman.HiddenWord(word))
	}
	win = hangman.CompareLetters([]rune(hiddenword), []rune(word), a)
	if win == 1 {
		fmt.Println("\nCongrats !")
	} else if win == 0 {
		fmt.Println("\nGame Over !!! You're finished\n\nthe word was :", word)
	}
	fmt.Println("")
}
