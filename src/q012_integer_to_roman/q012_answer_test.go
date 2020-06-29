package q012_integer_to_roman

import (
	"strings"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	doTest(t, intToRoman)
}

func intToRoman(num int) string {
	sb := strings.Builder{}
	for i, n := range nums {
		for count := num/n - 1; count >= 0; count-- {
			sb.WriteByte(romans[i])
		}
		num = num % n
		if num >= limits[i] {
			// 左减不能跨过一个位数, 且不是5 的倍数
			minus := i + 2 - i%2
			num -= n - nums[minus]
			sb.WriteByte(romans[minus])
			sb.WriteByte(romans[i])
		}
	}
	return sb.String()
}

var nums = []int{1000, 500, 100, 50, 10, 5, 1}
var limits = []int{900, 400, 90, 40, 9, 4, 1}
var romans = []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}

func TestIntToRoman2(t *testing.T) {
	doTest(t, intToRoman)
}

// LeetCode 上更快的一种优化解法
func intToRoman2(num int) string {
	b := strings.Builder{}
	nums := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; num > 0; i++ {
		j := num / nums[i]
		for j > 0 {
			b.WriteString(romans[i])
			j--
		}
		num %= nums[i]
	}
	return b.String()
}
