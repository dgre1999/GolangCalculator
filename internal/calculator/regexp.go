package calculator

import (
	"regexp"
)

// Explanation of the following regex: Any token, which is a digit or any of the accepted operands, followed by 1 or more whitespace (to allow Strings.fields)
// repeated at least one time, and then ending on a final token that does not need to be followed by whitespace. Does not check for expression structure. Structure can be implemented later if time or lacking ideas
func expressionRegex() *regexp.Regexp {
	pattern := `^\s*((-?\d+(\.\d+)?|\+|\-|\*|\/|%|\^|add|plus|subtract|minus|multiply|times|divide|mod|power|\(|\))\s+)+(-?\d+(\.\d+)?|\+|\-|\*|\/|%|\^|add|plus|subtract|minus|multiply|times|divide|mod|power|\(|\))\s*$`
	return regexp.MustCompile(pattern)
}
