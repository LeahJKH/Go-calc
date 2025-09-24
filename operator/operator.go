package operator

import "strings"
func ValidateOperator(input string) string {
	i := strings.TrimSpace(input)
	if i == "+" || i == "-" || i == "*" || i == "/" {
		return i
	}
	return "%"
}
