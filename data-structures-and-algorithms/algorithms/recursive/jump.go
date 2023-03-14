package recursive

func Junp(n int) {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Jump(n-1) + Jump(n-2)
}
