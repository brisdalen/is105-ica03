package main

import(
	"./lineshift"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("NOT ENOUGH ARGUMENTS TO RUN")
	}
	s := lineshift.DetermineLineshifts(args[1])
	fmt.Println(s)
}


