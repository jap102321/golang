package recursion

func Recursion(number int) int {
	if number == 0 {
		return 1
	}
	return number * Recursion(number-1)
}
