package frequence

import (
	"bufio"
	"fmt"
	"github.com/Nosp1/Is-105/is105-ica03/fileversion"
	//"github.com/go-delve/delve/pkg/dwarf/line"
	"os"
	//"path/filepath"
	"sort"
)

/*
  oppgave 3
 hjelpefunksjon for å legge til kommandolinje argument med -f <filnavn>
*/
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
			path = args[2]
		}
	}
}

var path string
var file string = "​bfrequence_res​1"
var dir = "./frequence/frequenceresults"

// buffered filereader. Tar in filenavn, teller linjer med hjelpe funksjonen LinesinFilebuffered.
func Bfrequence(fileName string) {
	f, _ := os.Open(fileName)
	for index, line := range LinesInFileBuffered(fileName) {
		fmt.Printf("Index = %v, line = %v\n", index, line)
	}
	defer f.Close()
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
	fmt.Println("dette er de fem mest brukte runene: ")
	counter := 0
	for _, lf := range list {
		if counter >= 5 {
			break
		}
		fmt.Printf("%+q %7d\n", lf.string, lf.freq)
		counter++
	}
	// Filen finnes alt
	_, err := os.Stat(path)
	fmt.Println(err)
	if err == nil {
		WriteToFile(fileversion.DontOverrideFileversion(path), list, lines)
		fmt.Println("Finnes")
		// Filen finnes ikke
	}
	if err != nil {
		// Lager ny(?)
			WriteToFile(path, list, lines)
			fmt.Println("Finnes ikke, lager ny")

	}
}

func WriteToFile(filepath string, list lflist, lines []string) string {
	write, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer write.Close()

	w := bufio.NewWriter(write)
	fmt.Println("Writing to file")
	fmt.Fprintln(w, "Bfrequence resultat: ")
	fmt.Fprint(w, "De mest brukte filene er ")
	for i := 0; i < 5; i ++ {
		_, err = fmt.Fprintf(w, "\n%+q %7d\n", list[i].string, list[i].freq)
	}
	_, err = fmt.Fprintf(w, "\nAntall linjer: %v", len(lines))
	w.Flush()
	fmt.Println("Writing to file is complete. ")

	return "sucess!"
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

func LinesInFileBuffered(fileName string) []string {
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
