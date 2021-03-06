package q001_two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	doTest(t, twoSum)
}

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i, num := range nums {
		if idx, exists := indexes[target-num]; exists {
			return []int{idx, i}
		}
		indexes[num] = i
	}
	return nil
}
