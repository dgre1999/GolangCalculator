package calculator

import "testing"

func TestBasicCompute(t *testing.T) {
	c := NewBasic()

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
			t.Fatalf("unexpected error: %v", err)
		}
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}

func TestRPNCompute(t *testing.T) {
	c := NewRPN()

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
