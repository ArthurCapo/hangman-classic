package hangman

import "fmt"

func PrintWord(word string) {
	for x := range word {
		fmt.Print(string(word[x]))
		fmt.Print(" ")
	}
}
