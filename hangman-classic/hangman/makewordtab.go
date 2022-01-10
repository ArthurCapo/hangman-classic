package hangman

import (
	"io/ioutil"
	"log"
)

func MakeWordTab(path string) []string {
	var tab []string
	var tmp []rune
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	for i := range content {
		if content[i] != '\n' {
			tmp = append(tmp, rune(content[i]))
		} else {
			tab = append(tab, string(tmp))
			tmp = nil
		}
	}
	return tab
}
