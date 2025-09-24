package numbertwo

import (
	"strconv"
)


func ParseEntry(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return num
}
