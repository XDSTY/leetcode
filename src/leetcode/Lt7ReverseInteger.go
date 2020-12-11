package leetcode

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
输入: 123
输出: 321
输入: -123
输出: -321

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
func main() {
}

func reverse(x int) int {
	tag := 1
	if x < 0 {
		tag = -1
		x = -x
	}
	var res int64 = 0
	for x > 0 {
		if res == 0 && x%10 == 0 {
			x /= 10
			continue
		}
		res *= 10
		res += int64(x % 10)
		x /= 10
	}
	if res > int64(1<<31-1) {
		return 0
	}
	return tag * int(res)
}
