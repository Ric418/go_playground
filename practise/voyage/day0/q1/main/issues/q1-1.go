package issues

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Fib(n int) (res int) {
	switch {
	case n < 0:
		return n
	default:
		f := fibonacci()
		for i := 0; i < n-1; i++ {
			f()
		}
		return f()
	}
}


