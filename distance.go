package levenshtein

import "fmt"

func Distance(s1, s2 string) int {
	d := matrix(s1, s2)
	return d[len(s1)][len(s2)]
}

func Operations(s1, s2 string) []string {
	d := matrix(s1, s2)

	var ops []string
	for i, j := len(s1), len(s2); i > 0 || j > 0; {
		switch {
		case i == 0:
			ops = append(ops, fmt.Sprintf("add %c", s2[j-1]))
			j = j - 1
		case j == 0:
			ops = append(ops, fmt.Sprintf("delete %c", s1[i-1]))
			i = i - 1
		case s1[i-1] == s2[j-1] && d[i][j] == d[i-1][j-1]:
			ops = append(ops, fmt.Sprintf("keep %c", s2[j-1]))
			i, j = i-1, j-1
		case d[i-1][j] <= d[i][j-1] && d[i-1][j] <= d[i-1][j-1]:
			ops = append(ops, fmt.Sprintf("delete %c", s1[i-1]))
			i = i - 1
		case d[i][j-1] <= d[i-1][j] && d[i][j-1] <= d[i-1][j-1]:
			ops = append(ops, fmt.Sprintf("add %c", s2[j-1]))
			j = j - 1
		default:
			ops = append(ops, fmt.Sprintf("swap %c for %c", s2[j-1], s1[i-1]))
			i, j = i-1, j-1
		}
	}

	for i, j := 0, len(ops)-1; i < j; i, j = i+1, j-1 {
		ops[i], ops[j] = ops[j], ops[i]
	}

	return ops
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
