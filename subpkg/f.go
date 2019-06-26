package subpkg

func Factorial(n int64) int64 {
	if n == 0 {
		return returnNumber(1)
	} else {
		return returnNumber(n) * Factorial(returnNumber(n-1))
	}
}

//go:noinline
func returnNumber(n int64) int64 {
	return n
}

