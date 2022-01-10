package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

func HiddenWord(word string) []rune {
	n := len(word)/2 - 1
	hiddenword := []rune(word)
	rword := []rune(word)
	var letter rune
	for i := range word {
		hiddenword[i] = 95
	}
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().Unix())
		rand := rand.Intn(len(word))
		letter = rword[rand]
		for j := range word {
			if letter == rword[j] {
				hiddenword[j] = letter
			}
		}
	}
	PrintWord(string(hiddenword))
	fmt.Println("")
	return hiddenword
}

func verif(word []rune) bool {
	for i := range word {
		if word[i] == '_' {
			return false
		}
	}
	return true
}

func CompareLetters(hiddenword []rune, word []rune, a int) int {
	var j int
	po := &a
	rma := false
	var m []string
	for ; a > 0; a-- {
		if rma {
			remains(a, false)
			rma = false
		}
		j = 0
		if verif(hiddenword) {
			return 1
		}
		letter := ""
		fmt.Print("\n\nLetter : ")
		t := time.NewTicker(time.Second * 10)
		go timer(t, po)
		fmt.Scanf("%s", &letter)
		t.Stop()
		if letter == "" {
			fmt.Println("\nYou don't enter a letter or a word")
			a++
			continue
		}
		if len(letter) != len(word) && len(letter) != 1 && letter != "pw" && letter != "STOP" {
			a -= 2
			remains(a, false)
			continue
		}
		for _, i := range m {
			if i == letter {
				fmt.Println("You use the same letter twice.")
				continue
			}
		}
		m = append(m, letter)
		if letter == "STOP" {
			save(hiddenword, word, a)
			fmt.Println("\nGame save in save.txt.")
			return 2
		} else if letter == "pw" {
			PrintWord(string(hiddenword))
			a++
			continue
		}
		if len(letter) != 1 && letter != "pw" && len(letter) != len(hiddenword) {
			a++
			fmt.Printf("\nOne single letter only")
			continue
		}
		fmt.Println("")
		rletter := []rune(letter)
		if len(rletter) == len(hiddenword) {
			boo := false
			for i := 0; i < len(hiddenword); i++ {
				if word[i] != rletter[i] {
					a -= 2
					remains(a, false)
					break
				} else if word[i] == rletter[i] && i != len(letter)-1 {
					boo = true
					continue
				} else if word[i] == rletter[i] && boo {
					return 1
				}
			}
			a++
			continue
		}
		for i := range hiddenword {
			if word[i] == rletter[0] {
				hiddenword[i] = rletter[0]
				j++
			}
		}
		if j > 0 {
			PrintWord(string(hiddenword))
			a++
			continue
		}
		rma = true
	}
	remains(0, false)
	return 0
}
