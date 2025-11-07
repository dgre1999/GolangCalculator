package calculator

import (
	"fmt"
)

type Calculator struct {
	history []string
}

func New() *Calculator {
	return &Calculator{history: []string{}}
}

func (c *Calculator) Compute(x, y float64, op string) (float64, error) {
	var result float64
	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result = x / y
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
	expr := fmt.Sprintf("%f %s %f = %f", x, op, y, result)
	c.history = append(c.history, expr)
	return result, nil
}

func (c *Calculator) History() []string {
	return c.history
}
