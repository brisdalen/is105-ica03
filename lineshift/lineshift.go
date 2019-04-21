package lineshift

import (
	"../fileutils"
	"path/filepath"
)

var cr bool
var lf bool

func DetermineLineshifts(filename string) string {
	b := fileutils.FileToByteslice(filename)
	cr = ContainsCR(b)
	lf = ContainsLF(b)

	if cr && lf {
		return cleanFilePath(filename) + " contains CRLF."
	}
	if !cr && lf {
		return cleanFilePath(filename) + " contains LF only."
	}

	return "Neither LF or CRLF"
}

func ContainsCR(input []byte) bool {
	return contains(input, 13)
}

func ContainsLF(input []byte) bool {
	return contains(input, 10)

}
// Hjelpefunksjon for å sjekke om en gitt slice har en karakter sin byte-kode
func contains(slice []byte, c byte) bool {
	for _, b := range slice {
		if b == c {
			return true
		}
	}
	return false
}
// Hjelpefunksjon for å rengjøre filepathen i DetermineLineShifts
func cleanFilePath(filename string) string {
	return filepath.Base(filename)
}

