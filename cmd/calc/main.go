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
		log.Println("Failed to load icon:", err)
	} else {
		myApp.SetIcon(icon)
	}

	c := calc.NewCalculator()
	c.LoadUI(myApp)

	myApp.Lifecycle().SetOnStopped(func() {
		log.Println("Cleanup tasks completed. Application stopped.")
	})

	log.Println("Starting the application...")
	myApp.Run()
}
