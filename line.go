package breakline

import "strings"

// BreakLine returns a sequence of '='.
func BreakLine(n int) string {
	return strings.Repeat("=", n)
}
