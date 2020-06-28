package q007_reverse_integer

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	doTest(t, reverse)
}

func reverse(x int) int {
	var res int64
	num := int64(x)
	if x < 0 {
		num = -num
	}
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	if x < 0 {
		res = -res
	}
	if res < math.MinInt32 || math.MaxInt32 < res {
		return 0
	}
	return int(res)
}
