package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type adaptiveTextTheme struct {
	fyne.Theme
	window  fyne.Window
	factor  float32
	maxSize *int
}

func NewAdaptiveTextTheme(t fyne.Theme, w fyne.Window, factor float32, maxSize *int) *adaptiveTextTheme {
	return &adaptiveTextTheme{Theme: t, window: w, factor: factor, maxSize: maxSize}
}

func (t *adaptiveTextTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		var h float32 = 380

		if t.window != nil && t.window.Content() != nil {
			currentH := t.window.Content().Size().Height
			if currentH > 10 {
				h = currentH
			}
		}

		res := h * t.factor

		if res < 14 {
			return 14
		}
		if t.maxSize != nil && res > float32(*t.maxSize) {
			return float32(*t.maxSize)
		} else if res > 45 {
			return 45
		}
		return res
	}
	return t.Theme.Size(name)
}
