# levenshtein

[![Go Report Card](https://goreportcard.com/badge/github.com/nathanjcochran/levenshtein)](https://goreportcard.com/report/github.com/nathanjcochran/levenshtein)
[![GoDoc](https://godoc.org/github.com/nathanjcochran/levenshtein?status.svg)](https://godoc.org/github.com/nathanjcochran/levenshtein) 

This package provides functions for calculating the [Levenshtein
distance](https://en.wikipedia.org/wiki/Levenshtein_distance) (a type of
[edit distance](https://en.wikipedia.org/wiki/Edit_distance)) between two
strings, and for generating a minimal list of edit operations required to
convert the source string into the target string. It does this by building
an edit matrix according to the [Wagner-Fischer Algorithm](https://en.wikipedia.org/wiki/Wagner–Fischer_algorithm).
Alternative insertion/removal/swap costs can be provided as options. The list
of edit operations is retrieved via a recursive algorithm which reads off a
backtrace of edit operations from the matrix.

Documentation and examples can be found at [godoc.org](https://godoc.org/github.com/nathanjcochran/levenshtein)
