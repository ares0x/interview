package divide_conquer

// 求 x 的 n 次方

// 暴力求解
func myPowUgly(x float64, n int) float64 {
	result := float64(1)
	for i := 0; i < n; i++ {
		result = result * x
	}
	return result
}

func myPow(x float64, n int) float64 {

}
