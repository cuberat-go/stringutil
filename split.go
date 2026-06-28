package stringutil

import (
	// Built-in/core modules.
	"strings"
)

// Splits the input string on any of the strings in the `delims` and returns
// a slice of the resulting substrings. If the input string is empty, returns an
// empty slice. If the input string does not contain any of the strings in
// the `delims`, returns a slice containing the input string as the only
// element. If mergeDelim is true, multiple consecutive strings found from
// the `delims` are treated as a single delimiter.
func SplitOnDelims(s string, delims []string, mergeDelim bool) []string {
	parts := []string{s}

	for _, d := range delims {
		var newParts []string
		for _, part := range parts {
			splitParts := strings.Split(part, d)
			if mergeDelim {
				for _, sp := range splitParts {
					if sp != "" {
						newParts = append(newParts, sp)
					}
				}
			} else {
				newParts = append(newParts, splitParts...)
			}
		}
		parts = newParts
	}

	return parts
}

// Splits the input string on any whitespace characters and returns a slice of
// the resulting substrings. If the input string is empty, returns an empty
// slice. If the input string does not contain any whitespace characters,
// returns a slice containing the input string as the only element. If
// mergeDelim is true, multiple consecutive whitespace characters are treated
// as a single delimiter.
//
// The whitespace characters considered are: space, tab, newline, vertical tab,
// form feed, carriage return, non-breaking space, and next line.
func SplitOnSpace(s string, mergeDelim bool) []string {
	delims := []string{" ", "\t", "\n", "\v", "\f", "\r", "\u0085",
		"\u00a0"}
	return SplitOnDelims(s, delims, mergeDelim)
}
