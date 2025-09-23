package operator

import "fmt"

func GetOperator() string {
	var op string
	for {
		fmt.Print("Insert operator (+,-,*,/): ")
		// checks for the right operator and runs again if its not implimented
		fmt.Scan(&op)
		if op == "+" || op == "-" || op == "*" || op == "/" {
			return op
		}
		fmt.Println("Invalid operator, try again.")
	}
}
