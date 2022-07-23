package _09_fib_number

func fib(n int) int {
	if n < 2 {
		return n
	}

	a := 0
	b := 0
	t := 1 // 这里以及初始化了n等于2值
	for i := 1; i < n; i++ {
		a = b
		b = t
		t = a + b
	}
	return t
}
