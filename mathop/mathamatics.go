package mathop

import "fmt"

func Calculate(num1 int, op string, num2 int) {
	// does the math
	switch op {
	case "+":
		fmt.Println("Sum:", num1+num2)
	case "-":
		fmt.Println("Sum:", num1-num2)
	case "*":
		fmt.Println("Sum:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cant divide by zero, silly :3")
		} else {
			fmt.Println("Sum:", num1/num2)
		}
	default:
		fmt.Println("Uh oh. You wrote something wrong or its my bad code lol")
	}
}
