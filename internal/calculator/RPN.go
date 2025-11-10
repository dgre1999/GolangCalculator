package calculator

import (
	"fmt"
	"math"
	"strconv"
)

var precedence = map[string]int{
	"+": 1, "plus": 1, "add": 1,
	"-": 1, "minus": 1, "subtract": 1,
	"*": 2, "times": 2, "multiply": 2,
	"/": 2, "divide": 2,
	"%": 2, "mod": 2,
	"^": 3, "power": 3,
}

var rightAssoc = map[string]bool{
	"^": true,
}

func infixToPostfix(inputs []string) []string {
	var output []string
	var stack []string

	for _, input := range inputs {
		if isNumber(input) {

			output = append(output, input)
		} else if input == "(" {
			stack = append(stack, input)
		} else if input == ")" {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if top == "(" {
					stack = stack[:len(stack)-1]
					break
				} else {
					if rightAssoc[input] {
						if precedence[top] > precedence[input] {
							output = append(output, top)
							stack = stack[:len(stack)-1]
						} else {
							break
						}
					} else if precedence[top] >= precedence[input] {
						output = append(output, top)
						stack = stack[:len(stack)-1]
					} else {
						break
					}
				}
			}
		} else {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if top == "(" {
					break
				}
				if rightAssoc[input] {
					if precedence[top] > precedence[input] {
						output = append(output, top)
						stack = stack[:len(stack)-1]
					} else {
						break
					}
				} else if precedence[top] >= precedence[input] {
					output = append(output, top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, input)
		}
	}
	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output
}

func evalPostfix(inputs []string) float64 {
	var stack []float64
	operandLookup := initMap()
	for _, input := range inputs {
		if isNumber(input) {
			n, _ := strconv.ParseFloat(input, 64)
			stack = append(stack, n)
		} else {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch operandLookup[input] {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					fmt.Println("Error: Division by zero")
					return 0
				} else {
					stack = append(stack, a/b)
				}
			case "%":
				if b == 0 {
					fmt.Println("Error: Division by zero")
					return 0
				} else {
					stack = append(stack, math.Mod(a, b))
				}
			case "^":
				stack = append(stack, math.Pow(a, b))
			}
		}

	}
	return stack[0]
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
