package mymath

func Mypower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func MyAbs(number int) int {
	if number < 0 { // 음수이면
		return number * -1
	}
	return number
}
