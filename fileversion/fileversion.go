package fileversion

import (
	"fmt"
	"strconv"
	"unicode"
)

/**
	Lager nytt filnavn med nytt nummer.
	Funksjonen antar at .txt er i bruk.
	For eksempel: lang1.txt -> lang2.txt
	Ledende 0'er blir fjernet: lang01.txt -> lang2.txt
 */
func DontOverrideFileversion(filename string) string {
	s := filename[:len(filename)-4]

	var id int

	for i := len(s); i > 0; i-- {
		if unicode.IsLetter(rune(s[i-1])) {
			id, _ = strconv.Atoi(s[i:])
			s = s[:i]
			break
		}
	}

	id++
	s = s + strconv.Itoa(id) + ".txt"
	fmt.Println(s)
	return s

}
