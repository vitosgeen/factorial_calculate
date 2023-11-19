package model

type Data struct {
	A int
	B int
}

func CalculateFactorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= number; i++ {
		result *= i
	}
	return result
}
