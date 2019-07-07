package levenshtein

import "fmt"

type OpType int8

const (
	Insert OpType = iota
	Remove
	Keep
	Swap
)

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
		panic(fmt.Sprintf("invalid OpType: %d", o))
	}
}

type Operation struct {
	Type   OpType
	Index  int
	Char   byte
	Result string
}

func (o Operation) String() string {
	return fmt.Sprintf("%6s %c at index %d: %s", o.Type, o.Char, o.Index, o.Result)
}

type Matrix struct {
	s1     string
	s2     string
	matrix [][]int
}

func Build(s1, s2 string) *Matrix {
	m := make([][]int, len(s1)+1)
	for i := range m {
		m[i] = make([]int, len(s2)+1)
	}

	// Deletions to get to empty target string from input string
	for i := 1; i <= len(s1); i++ {
		m[i][0] = i
	}

	// Insertions to get to target string from empty string
	for j := 1; j <= len(s2); j++ {
		m[0][j] = j
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}

			m[i][j] = min(
				m[i-1][j]+1,
				m[i][j-1]+1,
				m[i-1][j-1]+cost,
			)
		}
	}

	return &Matrix{
		s1:     s1,
		s2:     s2,
		matrix: m,
	}
}

func Distance(s1, s2 string) int {
	return Build(s1, s2).Distance()
}

func Operations(s1, s2 string) []Operation {
	return Build(s1, s2).Operations()
}

func (m *Matrix) Distance() int {
	return m.matrix[len(m.s1)][len(m.s2)]
}

func (m *Matrix) Operations() []Operation {
	ops := m.backtrace(len(m.s1), len(m.s2))
	return ops[1:]
}

func (m *Matrix) backtrace(i, j int) []Operation {
	switch {
	case j > 0 && m.matrix[i][j-1]+1 == m.matrix[i][j]:
		ops := m.backtrace(i, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Insert,
			Index:  j - 1,
			Char:   m.s2[j-1],
			Result: prev.Result[:j-1] + m.s2[j-1:j] + prev.Result[j-1:],
		})
	case i > 0 && m.matrix[i-1][j]+1 == m.matrix[i][j]:
		ops := m.backtrace(i-1, j)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Remove,
			Index:  j,
			Char:   prev.Result[j],
			Result: prev.Result[:j] + prev.Result[j+1:],
		})
	case i > 0 && j > 0 && m.matrix[i-1][j-1]+1 == m.matrix[i][j]:
		ops := m.backtrace(i-1, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Swap,
			Index:  j - 1,
			Char:   m.s2[j-1],
			Result: prev.Result[:j-1] + m.s2[j-1:j] + prev.Result[j:],
		})
	case i > 0 && j > 0 && m.matrix[i-1][j-1] == m.matrix[i][j]:
		ops := m.backtrace(i-1, j-1)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Keep,
			Index:  j - 1,
			Char:   prev.Result[j-1],
			Result: prev.Result,
		})
	default:
		return []Operation{{Result: m.s1}}
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
