package leetcode

// 汉诺塔问题
/**
f(1) : 1  只需要将1这个盘子移动到最后一根即可
f(2) : f(1) + 1 + f(1)  需要将编号为1的盘子移动到中间，再将2号盘子移动到最后一根，再将中间的1号盘子移到最后一根
f(3) : f(2) + 1 + f(2)  同上，需要将编号为(1,2)的盘子移到中间，再将3号盘子移动到最后一根，再将中间的（1，2）号盘子移到最后一根
*/
func main() {
	A := []int{2, 1, 0}
	B := make([]int, 0)
	C := make([]int, 0)
	hanota(A, B, C)
}

func hanota(A []int, B []int, C []int) []int {
	hanotaRecr(&A, &B, &C, len(A))
	return C
}

func hanotaRecr(A *[]int, B *[]int, C *[]int, n int) {
	// 只有一个盘子的时候，将盘子从A移动到C
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		return
	}
	// 将n-1个盘子从A经过C移动B
	hanotaRecr(A, C, B, n-1)
	// 将剩下的盘子从A移动到C
	hanotaRecr(A, B, C, 1)
	// 将n-1个盘子从B移动到C
	hanotaRecr(B, A, C, n-1)
}
