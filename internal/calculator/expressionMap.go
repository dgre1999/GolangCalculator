package calculator

// Map of accepted operand strings to their symbols. Moved to separate file to avoid a bloated calculator file
func initMap() map[string]string {
	operandLookup := map[string]string{
		"+": "+", "add": "+", "plus": "+",
		"-": "-", "subtract": "-", "minus": "-",
		"*": "*", "multiply": "*", "times": "*",
		"/": "/", "divide": "/",
		"%": "%", "mod": "%",
		"^": "^", "power": "^",
	}
	return operandLookup
}
