package main

import (
	"sort"
)

type three struct {
	a, b, c int
}

func (t three) sort() three {
	tt := []int{t.a, t.b, t.c}
	sort.Ints(tt)
	return three{tt[0], tt[1], tt[2]}
}

func threeSum(nums []int) [][]int {
	threes := make(map[three]bool) // a unique set
	counts := count(nums)
	for _, n := range nums {
		for _, t := range findThrees(counts, n) {
			threes[t] = true
		}
	}

	result := make([][]int, 0, len(threes))
	for k, _ := range threes {
		result = append(result, []int{k.a, k.b, k.c})
	}
	return result
}

func findThrees(counts map[int]int, target int) []three {
	var threes []three
	for n, ncount := range counts {
		nn := (-target) - n
		nncount, ok := counts[nn]

		if !ok {
			continue
		}

		if n == target && ncount < 2 {
			continue
		}

		if nn == target && nncount < 2 {
			continue
		}

		if n == nn && ncount < 2 {
			continue
		}

		if n == target && nn == target && ncount < 3 {
			continue
		}

		t := three{target, n, nn}.sort()
		threes = append(threes, t)
	}

	return threes
}

func count(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, n := range nums {
		c, _ := counts[n]
		counts[n] = c + 1
	}
	return counts
}
