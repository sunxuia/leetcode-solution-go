package q018_4sum

import (
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	doTest(t, fourSum)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for left := 0; left < len(nums)-2; left++ {
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}
		for right := len(nums) - 1; left < right; right-- {
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				continue
			}
			i, j := left+1, right-1
			for i < j {
				sum := nums[left] + nums[i] + nums[j] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[left], nums[right], nums[j]})
					i++
					j--
					for i < j && nums[i-1] == nums[i] {
						i++
					}
					for i < j && nums[j] == nums[j+1] {
						j--
					}
				} else if sum < target {
					i++
				} else {
					j--
				}
			}
		}
	}
	return res
}
