package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
