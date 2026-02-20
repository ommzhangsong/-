package main

import "sort"

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]
		if last[1] >= cur[0] {
			last[1] = max(last[1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
