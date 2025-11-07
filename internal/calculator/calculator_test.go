package calculator

import "testing"

func TestCompute(t *testing.T) {
	c := New()

	tests := []struct {
		a, b   float64
		op     string
		expect float64
	}{
		{2, 3, "+", 5},
		{10, 5, "-", 5},
		{2, 4, "*", 8},
		{9, 3, "/", 3},
	}

	for _, tt := range tests {
		res, err := c.Compute(tt.a, tt.b, tt.op)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != tt.expect {
			t.Errorf("expected %v, got %v", tt.expect, res)
		}
	}
}
