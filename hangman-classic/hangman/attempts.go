package hangman

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func remains(a int, bol bool) {
	if a == 10 {
		return
	}
	if !bol {
		fmt.Printf("Not present in the word, %d attempts remaining\n\n", a)
		print_boy(a)
		return
	}
}

func print_boy(a int) {
	count := 0
	content, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(content) && string(content[i]) != strconv.Itoa(a); i++ {
		count++
	}
	count++
	for ; content[count] != '$' && count < 730; count++ {
		fmt.Printf("%s", string(content[count]))
	}
}
