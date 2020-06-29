package q015_3sum

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	doTest(t, threeSum)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[low] + nums[i] + nums[high]
			if sum == 0 {
				res = append(res, []int{nums[low], nums[i], nums[high]})
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if sum > 0 {
				high--
			} else {
				low++
			}
		}
	}
	return res
}
