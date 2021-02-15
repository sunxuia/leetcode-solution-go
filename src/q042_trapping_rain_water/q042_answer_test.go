package q042_trapping_rain_water

import (
	"testing"
)

func TestTrap(t *testing.T) {
	doTest(t, trap)
}

// 递减栈
func trap(height []int) int {
	res := 0
	stack := make([]int, 0)
	for i, right := range height {
		tail := len(stack) - 1
		for ; tail > 0 && height[stack[tail]] < right; tail-- {
			// left, bottom, right 组成了1 个蓄水池
			bottom := height[stack[tail]]
			left := height[stack[tail-1]]
			if left > right {
				res += (i - stack[tail-1] - 1) * (right - bottom)
			} else {
				res += (i - stack[tail-1] - 1) * (left - bottom)
			}
		}
		if tail == 0 && height[stack[tail]] < right {
			stack = stack[:tail]
		} else {
			stack = stack[:tail+1]
		}
		stack = append(stack, i)
	}
	return res
}
