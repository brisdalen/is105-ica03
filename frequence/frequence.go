package frequence

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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



func LinesInFile(fileName string) []string {
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

func Frequence(fileName string) {
	for index, line := range LinesInFile(fileName) {
		fmt.Printf("Index = %v, line = %v\n", index, line)
	}
	// Get count of lines.
	lines := LinesInFile(fileName)
	fmt.Println("Antall linjer: ", len(lines))
	//Teller runes
	m := make(map[string]int)
	input, _ := ioutil.ReadFile(fileName)
	for i := 0; i < len(input); i++ {
		m[string(input[i])] += 1
	}
	//Lager en struct med en key av typen string
	//og en value av typen int
	type kv struct {
		Key   string
		Value int
	}
	//Setter variabelen ss til å være en slice av typen k og v
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	//Sorterer slicen ss
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
