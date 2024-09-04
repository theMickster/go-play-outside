package calculator

import (
	"fmt"
	"strconv"
)

// Calculate is the process for executing a mathematical operation on two values.
func Calculate(input1 string, input2 string, operation string) float64 {
	leftInput := convertInputToValue(input1)
	rightInput := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(leftInput, rightInput)
	case "-":
		return subtractValues(leftInput, rightInput)
	case "*":
		return multiplyValues(leftInput, rightInput)
	case "/":
		return divideValues(leftInput, rightInput)
	default:
		fmt.Println("Please enter a valid mathematical operator")
		return 0
	}
}

func convertInputToValue(input string) float64 {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a valid number", input)
		panic(message)
	}
	return result
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(dividend, divisor float64) (result float64) {
	defer func() {
		if message := recover(); message != nil {
			result = 0
		}
	}()
	return dividend / divisor
}
