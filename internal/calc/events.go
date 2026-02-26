package calc

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
)

func (c *Calculator) setupEvents() {
	canvas := c.window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.SetOnTypedKey(c.onTypedKey)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)
}

func (c *Calculator) onTypedRune(r rune) {
	s := string(r)
	if s == "c" {
		s = "C"
	}

	if btn, ok := c.buttons[s]; ok {
		btn.OnTapped()
	}
}

func (c *Calculator) onTypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		c.evaluate()
	case fyne.KeyBackspace:
		c.backspace()
	}
}

func (c *Calculator) onPasteShortcut(shortcut fyne.Shortcut) {
	content := shortcut.(*fyne.ShortcutPaste).Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err != nil {
		return
	}

	c.display(append(c.expression, content))
}

func (c *Calculator) onCopyShortcut(shortcut fyne.Shortcut) {
	shortcut.(*fyne.ShortcutCopy).Clipboard.SetContent(strings.Join(c.expression, ""))
}
