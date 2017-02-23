// Package stringutil contains utility functions for working with strings
package stringutil

// Revers returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}
