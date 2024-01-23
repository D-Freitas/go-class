package function

func Soma(number1, number2 int) int {
	if number1 == 0 || number2 == 0 {
		return 0
	}
	return number1 + number2
}

func Substracao(number1, number2 int) int {
	if number1 == 0 || number2 == 0 {
		return 0
	}
	return number1 - number2
}
