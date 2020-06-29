package q008_string_to_integer_atoi

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	doTest(t, myAtoi)
}

func myAtoi(str string) int {
	n := len(str)
	i := 0
	for i < n && str[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	isNegative := str[i] == '-'
	var res int64
	if str[i] == '-' || str[i] == '+' {
		i++
	}
	for ; i < n && '0' <= str[i] && str[i] <= '9'; i++ {
		if isNegative {
			res = res*10 - int64(str[i]-'0')
			if res < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			res = res*10 + int64(str[i]-'0')
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}
	return int(res)
}
