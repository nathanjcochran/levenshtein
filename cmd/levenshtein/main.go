package main

import (
	"fmt"
	"os"

	"github.com/nathanjcochran/levenshtein"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s: missing required arguments\n", os.Args[0])
		usage()
	} else if len(os.Args) > 3 {
		fmt.Printf("%s: too many arguments provided\n", os.Args[0])
		usage()
	}

	matrix := levenshtein.Build(os.Args[1], os.Args[2])

	fmt.Printf("Edit distance: %d\n", matrix.Distance())
	fmt.Printf("Operations:\n")
	for _, op := range matrix.Operations() {
		fmt.Printf(" %s\n", op)
	}
}

func usage() {
	fmt.Printf("Usage: %s <source> <target>\n", os.Args[0])
	os.Exit(2)
}
