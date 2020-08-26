package q040_combination_sum_ii

import (
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	doTest(t, combinationSum2)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfs(&res, target, candidates, 0, make([]int, 0))
	return res
}

func dfs(res *[][]int, target int, candidates []int, index int, path []int) {
	if target == 0 {
		fixedPath := make([]int, len(path))
		copy(fixedPath, path)
		*res = append(*res, fixedPath)
		return
	}
	for i := index; i < len(candidates); i++ {
		// 去重
		if index < i && candidates[i-1] == candidates[i] {
			continue
		}
		if target < candidates[i] {
			break
		}
		path = append(path, candidates[i])
		dfs(res, target-candidates[i], candidates, i+1, path)
		path = path[:len(path)-1]
	}
}
