package frequence

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//  hjelpefunksjon for å legge til kommandolinje argument med -f <filnavn>
func HovedBfrequence(filename string) {
	args := os.Args
	// om kommandolinjearugment er for kort kommer feilmelding.
	if len(args) == 2 {
		feilmelding := "Du mangler -f"
		panic(feilmelding)
	}
	// om lengden på argumentene i kommandolinje er  3 og posisjon 1 har "-f" så kjører Befrequence.
	if len(args) == 3 {
		if args[1] == "-f" {
			Bfrequence(args[2])
		}
	}
}

// buffered filereader. Tar in filenavn, teller linjer med hjelpe funksjonen LinesinFilebuffered.
func Bfrequence(fileName string) {
	f, _ := os.Open(fileName)
	for index, line := range LinesInFileBuffered(fileName) {
		fmt.Printf("Index = %v, line = %v\n", index, line)
	}
	// Get count of lines.
	lines := LinesInFileBuffered(fileName)
	fmt.Println("Antall linjer: ", len(lines))

	//oppretter scanner
	input := bufio.NewScanner(f)
	//splitter scanna runer.
	input.Split(bufio.ScanRunes)
	//oppretter map av lest fil
	m := make(map[string]int)
	//scanner filen.
	for input.Scan() {
		ru := input.Text()
		//inkrementerer runer
		m[ru]++
	}
	//gjør variablel list til lflist og lengden på m.
	list := make(lflist, 0, len(m))
	// iterer over map
	for l, f := range m {
		//legger til freq og string til listen.
		list = append(list, &letterFreq{l, f})
	}
	// sorterer listen
	sort.Sort(list)
	fmt.Println("these are the five most used: ")
	counter := 0
	for _, lf := range list {
		if counter >= 5 {
			break
		}
		fmt.Printf("%+q %7d\n", lf.string, lf.freq)
		counter++
	}
}

type letterFreq struct {
	string
	freq int
}
type lflist []*letterFreq

func (list lflist) Len() int {
	return len(list)
}
func (list lflist) Less(i, j int) bool {
	switch t := list[i].freq - list[j].freq; {
	case t < 0:
		return false
	case t > 0:
		return true
	}
	return list[i].string < list[j].string
}
func (list lflist) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func LinesInFileBuffered(fileName string) ([]string) {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result

}
