package main

import (
	"sort"
)

/**
Given a collection of intervals, merge all overlapping intervals.
给一个集合的区间，合并有覆盖的区间
Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
*/

func main() {

}

type LL [][]int

func (arr LL) Len() int {
	return len(arr)
}

func (arr LL) Less(i, j int) bool {
	if arr[i][0] < arr[j][0] {
		return true
	} else if arr[i][0] > arr[j][0] {
		return false
	} else if arr[i][1] < arr[j][1] {
		return true
	}
	return false
}

func (arr LL) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	arr := LL(intervals)
	sort.Sort(arr)
	res := make([][]int, 0)
	tarr := make([]int, 2)
	copy(tarr, arr[0])
	res = append(res, tarr)
	for i := 1; i < len(arr); i++ {
		if res[len(res)-1][0] <= arr[i][0] && arr[i][0] <= res[len(res)-1][1] {
			if arr[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = arr[i][1]
			}
		} else {
			tarr := make([]int, 2)
			copy(tarr, arr[i])
			res = append(res, tarr)
		}
	}
	return res
}
