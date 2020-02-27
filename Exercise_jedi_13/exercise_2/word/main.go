// Package word provides function for strings
package word

import "strings"

// UseCount returns the number of times a word is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns total number of words in the given string
func Count(s string) int {
	xs := strings.Fields(s)
  return len(xs)
}
