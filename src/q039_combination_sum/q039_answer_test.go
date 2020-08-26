package q039_combination_sum

import (
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	doTest(t, combinationSum)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfs(&res, target, candidates, 0, []int{})
	return res
}

func dfs(res *[][]int, target int, candidates []int, index int, path []int) {
	if target == 0 {
		fixedPath := make([]int, len(path))
		copy(fixedPath, path)
		*res = append(*res, fixedPath)
		return
	}
	if index == len(candidates) {
		return
	}
	val := candidates[index]
	repeatTimes := target / val
	for i := 0; i <= repeatTimes; i++ {
		dfs(res, target-val*i, candidates, index+1, path)
		path = append(path, val)
	}
	path = path[:len(path)-repeatTimes-1]
}
