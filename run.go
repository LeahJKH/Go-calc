package main

import (
	"fmt"
	// importing of smaller modules
	"go-calc/mathop"
	"go-calc/numberone"
	"go-calc/numbertwo"
	"go-calc/operator"
)

func main() {
	//the big brains O_O
	fmt.Print("Welcome to Go-Calc! Press Ctrl+C to exit.\n")
	for {
		num1 := numberone.GetNumberOne()
		op := operator.GetOperator()
		num2 := numbertwo.GetNumberTwo()
		mathop.Calculate(num1, op, num2)
		fmt.Print("\n")
	}
}
