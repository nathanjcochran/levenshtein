package levenshtein_test

import (
	"fmt"

	"github.com/nathanjcochran/levenshtein"
)

func Example() {
	matrix := levenshtein.Build("horse", "arose")

	fmt.Printf("Matrix:\n%s\n\n", matrix)
	fmt.Printf("Edit distance: %d\n", matrix.Distance())
	fmt.Printf("Operations:\n")
	for _, op := range matrix.Operations() {
		fmt.Printf(" %s\n", op)
	}

	// Output:
	// Matrix:
	// 0 1 2 3 4 5
	// 1 1 2 3 4 5
	// 2 2 2 2 3 4
	// 3 3 2 3 3 4
	// 4 4 3 3 3 4
	// 5 5 4 4 4 3
	//
	// Edit distance: 3
	// Operations:
	//    swap a at index 0: aorse
	//  remove o at index 1: arse
	//    keep r at index 1: arse
	//  insert o at index 2: arose
	//    keep s at index 3: arose
	//    keep e at index 4: arose
}

func ExampleBuild() {
	fmt.Println(levenshtein.Build("horse", "arose"))

	// Output:
	// 0 1 2 3 4 5
	// 1 1 2 3 4 5
	// 2 2 2 2 3 4
	// 3 3 2 3 3 4
	// 4 4 3 3 3 4
	// 5 5 4 4 4 3
}

func ExampleMatrix_Distance() {
	matrix := levenshtein.Build("horse", "arose")
	fmt.Println(matrix.Distance())

	// Output:
	// 3
}

func ExampleDistance() {
	fmt.Println(levenshtein.Distance("horse", "arose"))

	// Output:
	// 3
}

func ExampleMatrix_Operations() {
	matrix := levenshtein.Build("horse", "arose")
	for _, op := range matrix.Operations() {
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
