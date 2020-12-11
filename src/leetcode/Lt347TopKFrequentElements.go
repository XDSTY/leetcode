package leetcode

import "fmt"

/**
Given a non-empty array of integers, return the k most frequent elements.
找出出现最频繁的k个数字
思路：先使用map统计每个数字的出现次数，再构建一个k大小最小堆
*/

func main() {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := numsMap[nums[i]]; ok {
			numsMap[nums[i]] = val + 1
		} else {
			numsMap[nums[i]] = 1
		}
	}
	headArr := make([]int, k)
	headArrCount := 0
	for key, val := range numsMap {
		if headArrCount < k {
			headArr[headArrCount] = key
			headArrCount++
			topKFrequentBuildMinHeap(headArr, numsMap, headArrCount-1)
		} else if numsMap[headArr[0]] < val {
			headArr[0] = key
			topKFrequentBuildDownHeap(headArr, numsMap, 0, k)
		}
	}
	return headArr
}

func topKFrequentBuildDownHeap(nums []int, numsMap map[int]int, idx int, length int) {
	minIdx := idx
	if idx*2+1 < length && numsMap[nums[minIdx]] > numsMap[nums[idx*2+1]] {
		minIdx = 2*idx + 1
	}
	if idx*2+2 < length && numsMap[nums[minIdx]] > numsMap[nums[idx*2+2]] {
		minIdx = 2*idx + 2
	}
	if idx != minIdx {
		topKFrequentSwap(nums, idx, minIdx)
		topKFrequentBuildDownHeap(nums, numsMap, minIdx, length)
	}
}

func topKFrequentBuildMinHeap(nums []int, numsMap map[int]int, idx int) {
	if idx == 0 {
		return
	}
	if numsMap[nums[idx]] < numsMap[nums[(idx-1)/2]] {
		topKFrequentSwap(nums, idx, (idx-1)/2)
		topKFrequentBuildMinHeap(nums, numsMap, (idx-1)/2)
	}
}

func topKFrequentSwap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
