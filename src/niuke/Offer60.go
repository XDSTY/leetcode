package niuke

import "fmt"

/**
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
*/

func dicesProbability(n int) []float64 {
	point := 1.0 / 6
	arr1 := make([]float64, 7)
	for i := 1; i <= 6; i++ {
		arr1[i] = point
	}
	for i := 2; i <= n; i++ {
		tmpArr := make([]float64, i*6+1)
		for j := i - 1; j < len(arr1); j++ {
			for z := 1; z <= 6; z++ {
				tmpArr[z+j] = tmpArr[z+j]*0.1 + arr1[j]*point
			}
		}
		arr1 = tmpArr
	}
	return arr1[n:]
}

func main() {
	fmt.Println(dicesProbability(3))
}
