package main

import (
	"fmt"
	"os"

	"github.com/nathanjcochran/levenshtein"
)

func main() {
	fmt.Println(levenshtein.Distance(os.Args[1], os.Args[2]))
	for _, op := range levenshtein.Operations(os.Args[1], os.Args[2]) {
		fmt.Println(op)
	}
}
