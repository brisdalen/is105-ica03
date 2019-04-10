package frequence

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)



func Hovedfrequence(fileName string) {
	args := os.Args

	if len(args) == 2 {
		feilmelding := "Du mangler -f"
		panic(feilmelding)
	}

	if len(args) == 3 {
		if args[1] == "-f" {
			Frequence(args[2])
		}
	}
}


/*
//func LinesInFile(fileName string) []string {
//	f, _ := os.Open(fileName)
//	// Create new Scanner.
//	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}
*/
func Frequence(fileName string) {
	//for index, line := range LinesInFile(fileName) {
	//	fmt.Printf("Index = %v, line = %v\n", index, line)
	//}
	// Get count of lines.
	//lines := LinesInFile(fileName)
	//fmt.Println("Antall linjer: ", len(lines))
	//Teller runes
	file, err := ioutil.ReadFile(fileName)
	if err !=nil{
		fmt.Println(err)
	}
	lines := strings.Split(string(file), "\n")
	//kutter ned filen
	lines = lines [:len(lines)-1]
	//printer ut alle linjer med index. kommenter ut for renere output.
	for i, line := range lines{
		fmt.Println(i,line)
	}
	//printer ut total antall linjer i filen
	fmt.Println(fileName, "inneholder", len(lines), "linjer ")
	//oppretter map
	m := make(map[string]int)
	fmt.Println("dette er de fem mest brukte runene:")
	//iterer over filen og teller hver rune.
	for i := 0; i < len(file); i++ {
		m[string(file[i])] += 1
	}
	//Lager en struct med en key av typen string
	//og en value av typen int
	type kv struct {
		Key   string
		Value int
	}
	//Setter variabelen ss til å være en slice av typen k og v
	var ss []kv
	//iterer over map og legger runene inn i ss med k som string verdi for rune og v for int verdi for rune.
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	//Sorterer slicen ss etter mest brukte
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	//Oppretter ny variabel "ts" som returnerer de
	//fem mest brukte "runes" i ss
	ts := ss[:5]
	//Looper over key-values i ts-slicen.
	for _, kv := range ts {
		//Printer ut de fem mest brukte symbolene og antall ganger brukt.
		fmt.Printf("%q, %d\n", kv.Key, kv.Value)
	}

}
