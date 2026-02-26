package calc

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	expression []string

	output  *widget.Label
	scroll  *container.Scroll
	buttons map[string]*widget.Button
	window  fyne.Window
}

func NewCalculator() *Calculator {
	return &Calculator{
		buttons:    make(map[string]*widget.Button),
		expression: []string{},
	}
}
