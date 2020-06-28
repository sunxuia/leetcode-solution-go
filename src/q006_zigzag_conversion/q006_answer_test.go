package q006_zigzag_conversion

import (
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	doTest(t, convert)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	round := 2*numRows - 2
	sb := strings.Builder{}
	sb.Grow(n)
	for r := 0; r < numRows; r++ {
		for i := r; i < n; i += round {
			sb.WriteByte(s[i])
			if 0 < r && r < numRows-1 {
				other := i + round - 2*r
				if other < n {
					sb.WriteByte(s[other])
				}
			}
		}
	}
	return sb.String()
}
