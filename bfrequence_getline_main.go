package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	for index, line := range LinesInFile("pg100.txt") {
		fmt.Printf("Index = %v, line = %v\n", index, line)
	}
	// Get count of lines.
	lines := LinesInFile("pg100.txt")
	fmt.Println(len(lines))

	m := make(map[string]int)
	input, _ := ioutil.ReadFile("pg100.txt")

	for i := 0; i < len(input); i++ {
		m[string(input[i])] += 1
	}
	fmt.Println(m)

	for key, value := range m {
		fmt.Printf("Key:%q", key, "Value:", value, "\n")
	}

}
