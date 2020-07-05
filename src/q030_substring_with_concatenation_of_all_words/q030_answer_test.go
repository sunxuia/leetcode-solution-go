package q030_substring_with_concatenation_of_all_words

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	doTest(t, findSubstring)
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}
	root := &nary{}
	for i, word := range words {
		node := root
		for _, c := range word {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &nary{}
			}
			node = node.children[c-'a']
		}
		if node.idxes == nil {
			node.idxes = []int{}
		}
		node.idxes = append(node.idxes, i)
	}

	size, width := len(words), len(words[0])
	res := []int{}
	for i, limit := 0, len(s)-size*width; i <= limit; i++ {
		exists := map[int]bool{}
		for j := 0; j < size; j++ {
			node := root
			for k := 0; k < width && node != nil; k++ {
				node = node.children[s[i+j*width+k]-'a']
			}
			if node == nil || node.idxes == nil {
				break
			} else {
				node.t = (node.t + 1) % len(node.idxes)
				if exists[node.idxes[node.t]] {
					break
				}
				exists[node.idxes[node.t]] = true
			}
		}
		if len(exists) == size {
			res = append(res, i)
		}
	}
	return res
}

type nary struct {
	t        int
	idxes    []int
	children [26]*nary
}
