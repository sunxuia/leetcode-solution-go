package q016_3sum_closest

import (
	"sort"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	doTest(t, threeSumClosest)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 20000
	for i := 1; i < len(nums)-1; i++ {
		left, right := 0, len(nums)-1
		for left < i && i < right {
			sum := nums[left] + nums[i] + nums[right]
			if diff(res, target) > diff(sum, target) {
				res = sum

				// (不加这个OJ 上比较慢)
				if res == target {
					return res
				}
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
