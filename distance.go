package levenshtein

import "fmt"

// Default costs for inserting, removing, and swapping characters.
const (
	DefaultInsertCost = 1
	DefaultRemoveCost = 1
	DefaultSwapCost   = 1
)

// OpType represents a type of edit operation.
type OpType int8

const (
	Insert OpType = iota
	Remove
	Keep
	Swap
)

// String returns the string representation of an operation type.
func (o OpType) String() string {
	switch o {
	case Insert:
		return "insert"
	case Remove:
		return "remove"
	case Keep:
		return "keep"
	case Swap:
		return "swap"
	default:
		return "invalid"
	}
}

// Operation represents one of the operations performed on a source string
// during the process of converting it into a target string. Contains
// information about the type of operation, the character affected, the index
// at which the operation occured, and the intermediate result of performing
// this operation.
type Operation struct {
	Type   OpType
	Char   rune
	Index  int
	Result string
}

// String returns the string representation of an operation.
func (o Operation) String() string {
	return fmt.Sprintf("%6s %c at index %d: %s", o.Type, o.Char, o.Index, o.Result)
}

// Matrix contains a two-dimensional matrix used for calculating edit
// distances between two strings, and for retrieving a minimal list of edit
// operations for converting the source string into the target string.
type Matrix struct {
	matrix     [][]int
	source     []rune
	target     []rune
	insertCost int
	removeCost int
	swapCost   int
}

// An Option which can be applied when generating an edit matrix or
// calculating the edit distance between two strings - e.g. setting
// a non-default insert/removal/swap cost.
type Option func(m *Matrix)

// SetInsertCost is an option which allows you to set a custom insertion cost
// to use when calculating edit distances. If this option is not provided,
// DefaultInsertCost is used instead.
func SetInsertCost(cost int) Option {
	return func(m *Matrix) {
		m.insertCost = cost
	}
}

// SetRemoveCost is an option which allows you to set a custom removal cost
// to use when calculating edit distances. If this option is not provided,
// DefaultRemoveCost is used instead.
func SetRemoveCost(cost int) Option {
	return func(m *Matrix) {
		m.removeCost = cost
	}
}

// SetSwapCost is an option which allows you to set a custom swap cost to use
// when calculating edit distances. If this option is not provided,
// DefaultSwapCost is used instead.
func SetSwapCost(cost int) Option {
	return func(m *Matrix) {
		m.swapCost = cost
	}
}

// Builds and fills a matrix which can be used to calculate the edit distance
// between the two strings, or to retrieve a list of edit operations required
// to transform the source string into the target string.
func Build(source, target string, options ...Option) *Matrix {
	s := []rune(source)
	t := []rune(target)
	m := &Matrix{
		matrix:     newMatrix(s, t),
		source:     s,
		target:     t,
		insertCost: DefaultInsertCost,
		removeCost: DefaultRemoveCost,
		swapCost:   DefaultSwapCost,
	}
	for _, option := range options {
		option(m)
	}

	m.fill()
	return m
}

func newMatrix(source, target []rune) [][]int {
	m := make([][]int, len(source)+1)
	for i := range m {
		m[i] = make([]int, len(target)+1)
	}
	return m
}

func (m *Matrix) fill() {
	// Deletions to get to empty target string from input string
	for i := 1; i <= len(m.source); i++ {
		m.matrix[i][0] = i
	}

	// Insertions to get to target string from empty string
	for j := 1; j <= len(m.target); j++ {
		m.matrix[0][j] = j
	}

	// Fill rest of matrix, using cheapest of three options for filling each
	// cell (insert a character, delete a character, or swap a character)
	for i := 1; i <= len(m.source); i++ {
		for j := 1; j <= len(m.target); j++ {
			swapCost := m.swapCost
			if m.source[i-1] == m.target[j-1] {
				swapCost = 0
			}

			m.matrix[i][j] = min(
				m.matrix[i][j-1]+m.insertCost,
				m.matrix[i-1][j]+m.removeCost,
				m.matrix[i-1][j-1]+swapCost,
			)
		}
	}
}

// Distance builds a matrix and returns the edit distance between the two
// strings - i.e. the minimum number of edits required to transform the source
// string into the target string.
func Distance(source, target string) int {
	return Build(source, target).Distance()
}

// Operations builds a matrix and returns a minimal list of edit operations
// required to transform the source string into the target string.
func Operations(source, target string) []Operation {
	return Build(source, target).Operations()
}

// Distance returns the edit distance between the two strings - i.e. the
// minimum number of edits required to transform the source string into the
// target string.
func (m *Matrix) Distance() int {
	return m.matrix[len(m.source)][len(m.target)]
}

// Operations returns a minimal list of edit operations required to transform
// the source string into the target string.
func (m *Matrix) Operations() []Operation {
	ops := m.backtrace(len(m.source), len(m.target))
	return ops[1:] // Remove dummy operation
}

func (m *Matrix) backtrace(i, j int) []Operation {
	switch {
	case j > 0 && m.matrix[i][j-1]+m.insertCost == m.matrix[i][j]:
		ops := m.backtrace(i, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Insert,
			Char:   m.target[j-1],
			Index:  j - 1,
			Result: string(prev.Result[:j-1]) + string(m.target[j-1:j]) + string(prev.Result[j-1:]),
		})
	case i > 0 && m.matrix[i-1][j]+m.removeCost == m.matrix[i][j]:
		ops := m.backtrace(i-1, j)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Remove,
			Char:   m.source[i-1],
			Index:  j,
			Result: string(prev.Result[:j]) + string(prev.Result[j+1:]),
		})
	case i > 0 && j > 0 && m.matrix[i-1][j-1]+m.swapCost == m.matrix[i][j]:
		ops := m.backtrace(i-1, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Swap,
			Char:   m.target[j-1],
			Index:  j - 1,
			Result: string(prev.Result[:j-1]) + string(m.target[j-1:j]) + string(prev.Result[j:]),
		})
	case i > 0 && j > 0 && m.matrix[i-1][j-1] == m.matrix[i][j]:
		ops := m.backtrace(i-1, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Keep,
			Char:   m.target[j-1],
			Index:  j - 1,
			Result: prev.Result,
		})
	default:
		// Base case: return the original source string. This dummy operation
		// is removed before the final list of operations is returned.
		return []Operation{
			{Result: string(m.source)},
		}
	}
}

func min(nums ...int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
