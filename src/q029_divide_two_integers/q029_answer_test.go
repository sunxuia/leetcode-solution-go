package q029_divide_two_integers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	doTest(t, divide)
}

func divide(dividend int, divisor int) int {
	dend, dsor := int64(dividend), int64(divisor)
	if dend < 0 {
		dend = -dend
	}
	if dsor < 0 {
		dsor = -dsor
	}
	if dend < dsor {
		return 0
	}

	multi := 0
	for dsor<<multi < dend {
		multi++
	}

	res, curr := 1<<multi, dsor<<multi
	for curr <= dend-dsor || dend < curr {
		multi--
		if curr <= dend-dsor {
			curr += dsor << multi
			res += 1 << multi
		} else {
			curr -= dsor << multi
			res -= 1 << multi
		}
	}

	if (dividend < 0 || divisor < 0) && !(dividend < 0 && divisor < 0) {
		res = -res
	}
	if res < math.MinInt32 {
		res = math.MinInt32
	}
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	return res
}
