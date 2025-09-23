package main

import "fmt"

func main() {
	var num1 int
	var op string
	var num2 int

    fmt.Print("insert number 1: ")
	fmt.Scan(&num1)
	fmt.Print("insert operator (+,-,*,/): ")
	fmt.Scan(&op)
	if op != "+" || op != "-" || op != "/" || op != "*" {
		fmt.Print("put in a valid operator! \n\n\n")
		main()
	}

	
	fmt.Print("insert number 2: ")
	fmt.Scan(&num2)
	switch op {
	case "+":
		fmt.Print(num1 + num2)
	case "-":
		fmt.Print(num1 - num2)	
	case "*":
		fmt.Print(num1 * num2)	
	case "/":
		fmt.Print(num1 / num2)	
	default: 
		fmt.Print("you have written something invalid try again \n\n\n")
		main()	
	}
	
}