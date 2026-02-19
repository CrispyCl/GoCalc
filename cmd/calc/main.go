package main

import (
	"gocalc/internal/calc"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	c := calc.NewCalculator()
	c.LoadUI(app)
	app.Run()
}
