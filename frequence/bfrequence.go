package frequence

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

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

func Bfrequence() {
	//var count int
	for index, line := range LinesInFile("pg100.txt") {
		fmt.Printf("Index = %v, line = %v\n", index, line)
		//count = index
	}
	//fmt.Printf("Number of lines = %d", count)
	// Get count of lines.
	lines := LinesInFile("pg100.txt")
	fmt.Println(len(lines))

	//Teller runes
	m := make(map[string]int)
	input, _ := ioutil.ReadFile("./pg100.txt")

	for i := 0; i < len(input); i++ {
		m[string(input[i])] += 1
	}
	//fmt.Println(m)

	//Sorterer runes
	//var keys []string
	//for k := range m {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)

	//for _, k := range keys {
	//	fmt.Println("Key:", k, "Value:", m[k])
	//}

	//Sorterer runes #2
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	ts := ss[:5]

	for _, kv := range ts {
		fmt.Printf("%q, %d\n", kv.Key, kv.Value)
	}

}
