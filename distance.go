package levenshtein

import "fmt"

type OpType int8

const (
	Add OpType = iota
	Remove
	Swap
)

func (o OpType) String() string {
	switch o {
	case Add:
		return "add"
	case Remove:
		return "remove"
	case Swap:
		return "swap"
	default:
		panic(fmt.Sprintf("unexpected OpType: %v", o))
	}
}

type Operation struct {
	Type   OpType
	Index  int
	Char   byte
	Result string
}

func (o Operation) String() string {
	return fmt.Sprintf("%s %c at index %d: %s", o.Type, o.Char, o.Index, o.Result)
}

func Distance(s1, s2 string) int {
	d := matrix(s1, s2)
	return d[len(s1)][len(s2)]
}

func Operations(s1, s2 string) []Operation {
	d := matrix(s1, s2)
	ops := operations(s1, s2, len(s1), len(s2), d)
	return ops[1:]
}

func operations(s1, s2 string, i, j int, d [][]int) []Operation {
	switch {
	case i > 0 && d[i-1][j]+1 == d[i][j]:
		ops := operations(s1, s2, i-1, j, d)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Remove,
			Index:  j,
			Char:   prev.Result[j],
			Result: prev.Result[:j] + prev.Result[j+1:],
		})
	case j > 0 && d[i][j-1]+1 == d[i][j]:
		ops := operations(s1, s2, i, j-1, d)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Add,
			Index:  j - 1,
			Char:   s2[j-1],
			Result: prev.Result[:j-1] + s2[j-1:j] + prev.Result[j-1:],
		})
	case i > 0 && j > 0 && d[i-1][j-1]+1 == d[i][j]:
		ops := operations(s1, s2, i-1, j-1, d)
		prev := ops[len(ops)-1]
		return append(ops, Operation{
			Type:   Swap,
			Index:  j - 1,
			Char:   s2[j-1],
			Result: prev.Result[:j-1] + s2[j-1:j] + prev.Result[j:],
		})
	case i > 0 && j > 0 && d[i-1][j-1] == d[i][j]:
		return operations(s1, s2, i-1, j-1, d)
	default:
		return []Operation{{
			Result: s1,
		}}
	}
}

func matrix(s1, s2 string) [][]int {
	d := make([][]int, len(s1)+1)
	for i := range d {
		d[i] = make([]int, len(s2)+1)
	}

	// Deletions to get to empty target string from input string
	for i := 1; i <= len(s1); i++ {
		d[i][0] = i
	}

	// Insertions to get to target string from empty string
	for j := 1; j <= len(s2); j++ {
		d[0][j] = j
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}

			d[i][j] = min(
				d[i-1][j]+1,
				d[i][j-1]+1,
				d[i-1][j-1]+cost,
			)
		}
	}

	return d
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
