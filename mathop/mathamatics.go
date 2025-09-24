package mathop

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func CalculateGUI(num1 int, op string, num2 int, resultLabel *widget.Label) {

	switch op {
	case "+":
		resultLabel.SetText(fmt.Sprintf("Sum: %d", num1+num2))
	case "-":
		resultLabel.SetText(fmt.Sprintf("Sum: %d", num1-num2))
	case "*":
		resultLabel.SetText(fmt.Sprintf("Sum: %d", num1*num2))
	case "/":
		if num2 == 0 {
			resultLabel.SetText("Cant divide by zero, silly :3")
		} else {
			resultLabel.SetText(fmt.Sprintf("Sum: %d", num1/num2))
		}
	default:
		resultLabel.SetText("Uh oh. You wrote something wrong or its my bad code lol")
	}
}
