package main

import (
	"log"

	"gocalc/internal/calc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		log.Println("Не удалось загрузить иконку:", err)
	} else {
		myApp.SetIcon(icon)
	}

	c := calc.NewCalculator()
	c.LoadUI(myApp)
	myApp.Run()
}
