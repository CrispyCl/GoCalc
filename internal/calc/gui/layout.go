package gui

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type topLayout struct{}

func NewTopLayout() *topLayout {
	return &topLayout{}
}

func (t *topLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	objects[0].Resize(fyne.NewSize(size.Width, size.Height*0.2))
	objects[0].Move(fyne.NewPos(0, 0))
}

func (t *topLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}

type adaptiveLayout struct {
	rows int
}

func NewAdaptiveLayout(rows int) *adaptiveLayout {
	return &adaptiveLayout{rows: rows}
}

func (l *adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}

	pad := theme.Padding()
	maxAvailableHeight := size.Height * 0.8

	usableHeight := maxAvailableHeight - (pad * float32(l.rows-1))
	cellH := usableHeight / float32(l.rows)

	maxH := (size.Width - (pad * 3)) / 4
	if cellH > maxH {
		cellH = maxH
	}

	cellH = float32(math.Floor(float64(cellH)))
	totalHeight := (cellH * float32(l.rows)) + (pad * float32(l.rows-1))

	yOffset := size.Height - totalHeight - pad

	for i, obj := range objects {
		posY := float32(math.Round(float64(yOffset + float32(i)*(cellH+pad))))

		obj.Resize(fyne.NewSize(size.Width, cellH))
		obj.Move(fyne.NewPos(0, posY))
	}
}

func (l *adaptiveLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(300, 320)
}
