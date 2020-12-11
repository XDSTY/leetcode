package leetcode

import (
	"fmt"
	"sort"
)

/**Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
给一个数组的string，将他们按照拥有相同的字母进行分类
Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:
Input: strs = [""]
Output: [[""]]
Example 3:
Input: strs = ["a"]
Output: [["a"]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
*/

func main() {
	fmt.Println(groupAnagrams([]string{""}))
}

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string]int, 0)
	res := make([][]string, 0)
	for _, str := range strs {
		bytes := bytes(str)
		sort.Sort(bytes)
		tstr := string(bytes)
		if val, ok := strMap[tstr]; ok {
			res[val] = append(res[val], str)
		} else {
			arr := make([]string, 0)
			arr = append(arr, str)
			res = append(res, arr)
			strMap[tstr] = len(res) - 1
		}
	}
	return res
}
