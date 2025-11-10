package calculator

import (
	"fmt"
	"strings"
)

type Calculator interface {
	Compute(expression string) (float64, error)
	History() []string
}

type BasicCalculator struct {
	history []string
}

func NewBasic() *BasicCalculator {
	return &BasicCalculator{history: []string{}}
}

type RPNCalculator struct {
	history []string
}

func NewRPN() *RPNCalculator {
	return &RPNCalculator{history: []string{}}
}

func (c *BasicCalculator) Compute(expression string) (float64, error) {
	var operands = strings.Fields(expression)
	result, nil := EvalExpression(operands)
	expr := fmt.Sprintf("%s %s %s = %f", operands[0], operands[1], operands[2], result)
	c.history = append(c.history, expr)
	return result, nil
}

func (c *BasicCalculator) History() []string {
	return c.history
}

func (c *RPNCalculator) Compute(expression string) (float64, error) {
	operands := strings.Fields(expression)
	postfix := infixToPostfix(operands)
	result := evalPostfix(postfix)

	expr := fmt.Sprintf("%s = %f, RPN: %s", expression, result, postfix)
	c.history = append(c.history, expr)
	return result, nil
}

func (c *RPNCalculator) History() []string {
	return c.history
}
