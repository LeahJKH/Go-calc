package numberone

import "fmt"

func GetNumberOne() int {
	// makes the first number
	var num1 int
	fmt.Print("Insert number 1: ")
	fmt.Scan(&num1)
	return num1
}
