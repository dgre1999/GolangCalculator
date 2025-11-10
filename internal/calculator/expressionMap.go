package calculator

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
