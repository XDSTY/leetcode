package leetcode

import "fmt"

/**
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明:
如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。

输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3

输入:
gas  = [2,3,4]
cost = [3,4,3]
输出: -1

思路：假设从i站出发可以到达j站，则i+1站到j站之间的都不能绕一圈
证明：假设i+1可以绕一圈则i+1可以到j+1，又因为i可以到i+1，则i也可以到j+1，与i只能到j矛盾，上述命题成立

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gas-station
*/
func canCompleteCircuit(gas []int, cost []int) int {
	i, j, sumGas := 0, 0, 0
	for i < len(gas) {
		if sumGas+gas[j]-cost[j] >= 0 {
			sumGas += gas[j] - cost[j]
			j = (j + 1) % len(gas)
			if i == j {
				return i
			}
		} else {
			if j < i {
				return -1
			}
			j++
			i = j
			sumGas = 0
		}
	}
	return -1
}

func main() {
	/**
	[5,1,2,3,4]
	[4,4,1,5,1]
	*/
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
