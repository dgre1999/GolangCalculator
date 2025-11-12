package calculator

import (
	"fmt"
	"strconv"
)

// Evaluates an expression for the BasicCalculator type. Moved to separate file to avoid a bloated calculator file
func EvalExpression(expression []string) (float64, error) {
	operandLookup := initMap()
	if len(expression) != 3 {
		return 0, fmt.Errorf("invalid expression format")
	}
	var x, xerr = strconv.ParseFloat(expression[0], 64)
	var y, yerr = strconv.ParseFloat(expression[2], 64)
	if xerr != nil || yerr != nil {
		return 0, fmt.Errorf("invalid values")
	}
	var op = expression[1]
	var result float64
	switch operandLookup[op] {
	case "+":
		result = Add(x, y)
	case "-":
		result = Subtract(x, y)
	case "*":
		result = Multiply(x, y)
	case "/":
		result, _ = Divide(x, y)
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
	return result, nil
}

// Basic arithmetic operations
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
