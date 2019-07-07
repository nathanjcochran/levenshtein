package levenshtein_test

import (
	"fmt"

	"github.com/nathanjcochran/levenshtein"
)

func Example() {
	matrix := levenshtein.Build("horse", "arose")

	fmt.Printf("Edit distance: %d\n", matrix.Distance())
	fmt.Printf("Operations:\n")
	for _, op := range matrix.Operations() {
		fmt.Printf(" %s\n", op)
	}

	// Output:
	// Edit distance: 3
	// Operations:
	//    swap a at index 0: aorse
	//  remove o at index 1: arse
	//    keep r at index 1: arse
	//  insert o at index 2: arose
	//    keep s at index 3: arose
	//    keep e at index 4: arose
}

func ExampleDistance() {
	fmt.Println(levenshtein.Distance("horse", "arose"))

	// Output:
	// 3
}

func ExampleOperations() {
	for _, op := range levenshtein.Operations("horse", "arose") {
		fmt.Println(op)
	}

	// Output:
	//   swap a at index 0: aorse
	// remove o at index 1: arse
	//   keep r at index 1: arse
	// insert o at index 2: arose
	//   keep s at index 3: arose
	//   keep e at index 4: arose
}
