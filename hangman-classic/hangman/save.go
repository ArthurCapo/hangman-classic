package hangman

import (
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func save(hiddenword []rune, word []rune, a int) bool {
	var tmp [][]byte
	tmp = append(tmp, []byte(string(hiddenword)), []byte(string(word)), []byte(strconv.Itoa(a)))
	var data []byte
	for i, j := 0, 0; i < 3 && j < len(tmp[i]); j++ {
		if j == len(tmp[i])-1 {
			data = append(data, tmp[i][j])
			data = append(data, '\n')
			j = -1
			i++
			continue
		}
		data = append(data, tmp[i][j])
	}
	f, err1 := os.Create("save.txt")
	check(err1)
	err := os.WriteFile("save.txt", data, 0644)
	check(err)
	defer f.Close()
	return true
}
