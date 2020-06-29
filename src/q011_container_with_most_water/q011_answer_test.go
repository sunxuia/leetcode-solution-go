package q011_container_with_most_water

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	doTest(t, maxArea)
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		var area int
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if res < area {
			res = area
		}
	}
	return res
}
