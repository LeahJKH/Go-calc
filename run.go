package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"go-calc/themeium"
)



func main() {
	a := app.New()
	a.Settings().SetTheme(&themeium.MyTheme{})

	w := a.NewWindow("Calculator")

	input := binding.NewString()
	entry := widget.NewEntryWithData(input)

	makeBtn := func(label string) *widget.Button {
		return widget.NewButton(label, func() {
			switch label {
			case "C":
				input.Set("")
			case "Ce":
				val, _  := input.Get()
				if len(val) > 0 {
					input.Set(val[:len(val)-1])
				}
			case "=":
				//mathium.stringSplitter(input.Get())
				
			default:
				val, _ := input.Get()
				input.Set(val + label)
			}
		})
	}
	buttons := container.NewGridWithColumns(5,
		makeBtn("Ce"), makeBtn("C"), makeBtn("("), makeBtn(")"),makeBtn("abs("),
		makeBtn("7"), makeBtn("8"), makeBtn("9"), makeBtn("/"), makeBtn("acos("),
		makeBtn("4"), makeBtn("5"), makeBtn("6"), makeBtn("*"), makeBtn("acosh("),
		makeBtn("1"), makeBtn("2"), makeBtn("3"), makeBtn("-"), makeBtn("asin("),
		makeBtn("0"), makeBtn("."), makeBtn("="), makeBtn("+"),makeBtn("asinh("),
		makeBtn("atan("), makeBtn("atanh("), makeBtn("ceil("), makeBtn("tanh("),
		makeBtn("cos("), makeBtn("cosh("), makeBtn("exp("), makeBtn("floor("),
		makeBtn("log("), makeBtn("log10("), makeBtn("round("), makeBtn("sign("),
		makeBtn("sin("), makeBtn("sinh("), makeBtn("sqrt("), makeBtn("tan("),		
	)

	content := container.NewVBox(entry, buttons)
	w.SetContent(content)
	w.ShowAndRun()
}
