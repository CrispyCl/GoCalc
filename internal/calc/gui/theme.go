package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type largeTextTheme struct {
	fyne.Theme
	window fyne.Window
	label  *widget.Label
}

func NewLargeTextTheme(t fyne.Theme, w fyne.Window, l *widget.Label) *largeTextTheme {
	return &largeTextTheme{Theme: t, window: w, label: l}
}

func (t *largeTextTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		if t.window == nil || t.window.Content() == nil {
			return 28
		}

		h := t.window.Content().Size().Height
		if h <= 0 {
			return 28
		}

		res := h * 0.07

		if res < 24 {
			return 24
		}
		if res > 45 {
			return 45
		}
		return res
	}
	return t.Theme.Size(name)
}
