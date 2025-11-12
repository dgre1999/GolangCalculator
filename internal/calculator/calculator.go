package calculator

import (
	"fmt"
	"strings"
)

// Calculator interface, to allow for multiple types of calculators
type Calculator interface {
	Compute(expression string) (float64, error)
	History() []string
}

// The basic calculator. Allows for expressions of the form "x op y"
type BasicCalculator struct {
	history []string
}

// The RPN calculator. Allows for expressions of any length, as long as the operands used are any of "+", "-", "*", "/", "%", "^"
// Also supports parentheses for order of operations. This is all done by converting the expression to Reverse Polish Notation
type RPNCalculator struct {
	history []string
}

// Creates a new calculator based on the string provided, to avoid creating a constructor for each calculator type
func NewCalculator(version string) (Calculator, error) { // Factory-esque method?
	switch version {
	case "basic":
		return &BasicCalculator{history: []string{}}, nil
	case "rpn":
		return &RPNCalculator{history: []string{}}, nil
	default:
		return nil, fmt.Errorf("unsupported calculator version: %s", version)
	}
}

// The compute function for the BasicCalculator type. Checks if the regex is valid, splits the expression into operands (split based on whitespace), evaluates it and adds it to the history
func (c *BasicCalculator) Compute(expression string) (float64, error) {
	calcRegex := expressionRegex()
	if !calcRegex.MatchString(expression) {
		return 0, fmt.Errorf("invalid expression format")
	}
	var operands = strings.Fields(expression)
	result, err := EvalExpression(operands)
	if err != nil {
		return 0, err
	}
	expr := fmt.Sprintf("%s %s %s = %f", operands[0], operands[1], operands[2], result)
	c.history = append(c.history, expr)
	return result, nil
}

// Returns the history of a BasicCalculator
func (c *BasicCalculator) History() []string {
	return c.history
}

// The compute function for the RPNCalculator type. Checks if the regex is valid, splits the expression into operands (split based on whitespace),
//
//	converts it to Reverse Polish Notation, evaluates it and adds it to the history
func (c *RPNCalculator) Compute(expression string) (float64, error) {
	calcRegex := expressionRegex()
	if !calcRegex.MatchString(expression) {
		return 0, fmt.Errorf("invalid expression format")
	}
	operands := strings.Fields(expression)
	postfix := infixToPostfix(operands)
	result := evalPostfix(postfix)

	expr := fmt.Sprintf("%s = %f, RPN: %s", expression, result, postfix)
	c.history = append(c.history, expr)
	return result, nil
}

// Returns the history of an RPNCalculator
func (c *RPNCalculator) History() []string {
	return c.history
}
