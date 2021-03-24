package loop_sum

func LoopSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}
