package numbertwo

import (
	"fmt"
)

func GetNumberTwo() int {
	// makes second number
	var num2 int
	fmt.Print("Insert number 2: ")
	fmt.Scan(&num2)
	return num2
}
