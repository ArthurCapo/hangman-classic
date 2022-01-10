package hangman

import (
	"fmt"
	"time"
)

func timer(t *time.Ticker, a *int) {
	for range t.C {
		*a--
		fmt.Printf("You took too long time!!! %d attempts remaining\n\n", *a)
		print_boy(*a)
		fmt.Print("\n\nLetter : ")
		if *a <= 0 {
			return
		}
	}
}
