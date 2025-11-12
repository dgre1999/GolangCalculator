package calculator

import "testing"

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		expression []string
		expect     float64
	}{
		{[]string{"2", "plus", "3"}, 5},
		{[]string{"10", "-", "5"}, 5},
		{[]string{"2", "multiply", "4"}, 8},
		{[]string{"9", "/", "3"}, 3},
	}

	for _, tt := range tests {
		res, err := EvalExpression(tt.expression)
		if err != nil {
			t.Fatalf("unexpected error: %v %v", err, tt.expression)
		}
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}

func TestBasicCompute(t *testing.T) {
	c, _ := NewCalculator("basic")

	tests := []struct {
		expression string
		expect     float64
	}{
		{"2 plus 3", 5},
		{"10 - 5", 5},
		{"2 multiply 4", 8},
		{"9 / 3", 3},
	}

	for _, tt := range tests {
		res, err := c.Compute(tt.expression)
		if err != nil {
			t.Fatalf("unexpected error: %v %v", err, tt.expression)
		}
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}

func TestInfixToPostfic(t *testing.T) {
	tests := []struct {
		expression []string
		expect     []string
	}{
		{[]string{"3", "plus", "4", "-", "5"}, []string{"3", "4", "plus", "5", "-"}},
		{[]string{"10", "*", "(", "2", "+", "3", ")", "^", "2"}, []string{"10", "2", "3", "+", "2", "^", "*"}},
		{[]string{"(", "500", "/", "2", ")", "/", "(", "(", "5", "+", "5", ")", "divide", "2", ")"}, []string{"500", "2", "/", "5", "5", "+", "2", "divide", "/"}},
		{[]string{"100", "%", "30", "add", "4", "*", "2"}, []string{"100", "30", "%", "4", "2", "*", "add"}},
	}

	for _, tt := range tests {
		res := infixToPostfix(tt.expression)
		if len(res) != len(tt.expect) {
			t.Errorf("expected %v, got %v", tt.expect, res)
			continue
		}
		for i := range res {
			if res[i] != tt.expect[i] {
				t.Errorf("expected %v, got %v", tt.expect, res)
				break
			}
		}
	}
}

func TestEvalPostfix(t *testing.T) {
	tests := []struct {
		expression []string
		expect     float64
	}{
		{[]string{"3", "4", "plus", "5", "-"}, 2},
		{[]string{"10", "2", "3", "+", "2", "^", "*"}, 250},
		{[]string{"500", "2", "/", "5", "5", "+", "2", "divide", "/"}, 50},
		{[]string{"100", "30", "%", "4", "2", "*", "add"}, 18},
	}

	for _, tt := range tests {
		res := evalPostfix(tt.expression)
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}

func TestRPNCompute(t *testing.T) {
	c, _ := NewCalculator("rpn")

	tests := []struct {
		expression string
		expect     float64
	}{
		{"3 plus 4 - 5", 2},
		{"10 * ( 2 + 3 ) ^ 2", 250},
		{"( 500 / 2 ) / ( ( 5 + 5 ) divide 2 )", 50},
		{"100 % 30 add 4 * 2", 18},
	}

	for _, tt := range tests {
		res, err := c.Compute(tt.expression)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}

func TestRegex(t *testing.T) {
	calcRegex := expressionRegex()
	validExpressions := []string{
		"2 plus 3",
		"10 - 5",
		"2 multiply 4",
		"9 / 3",
		"( 3 + 4 ) * 5",
		"100 % 30 add 4 * 2",
	}
	invalidExpressions := []string{
		"2 ++ 3",
		"10 -- 5",
		"9 / zero",
	}

	for _, expr := range validExpressions {
		if !calcRegex.MatchString(expr) {
			t.Errorf("expected valid expression to match: %s", expr)
		}
	}

	for _, expr := range invalidExpressions {
		if calcRegex.MatchString(expr) {
			t.Errorf("expected invalid expression to not match: %s", expr)
		}
	}
}
