package main

import (
	"go-calc/mathop"
	"go-calc/numberone"
	"go-calc/numbertwo"
	"go-calc/operator"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go-Calc GUI")


	w.SetIcon(resourceIconPng)

	w.Resize(fyne.NewSize(400, 300))


	num1Entry := widget.NewEntry()
	num1Entry.SetPlaceHolder("Number 1")

	num2Entry := widget.NewEntry()
	num2Entry.SetPlaceHolder("Number 2")

	opEntry := widget.NewEntry()
	opEntry.SetPlaceHolder("Operator (+, -, *, /)")

	resultLabel := widget.NewLabel("Result will appear here")

	calcBtn := widget.NewButton("Calculate", func() {
		num1 := numberone.ParseEntry(num1Entry.Text)
		num2 := numbertwo.ParseEntry(num2Entry.Text)
		op := operator.ValidateOperator(opEntry.Text)

		mathop.CalculateGUI(num1, op, num2, resultLabel)
	})


	w.SetContent(container.NewVBox(
		num1Entry,
		opEntry,
		num2Entry,
		calcBtn,
		resultLabel,
	))

	w.ShowAndRun()
}
