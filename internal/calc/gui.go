package calc

import (
	"gocalc/internal/calc/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	screenHeight = float32(380)
	screenWeight = float32(300)
)

func (c *Calculator) LoadUI(app fyne.App) {
	c.output = &widget.Label{
		Alignment:  fyne.TextAlignTrailing,
		Truncation: fyne.TextTruncateOff,
	}
	c.output.TextStyle = fyne.TextStyle{Monospace: true}

	scrollContainer := container.NewHScroll(c.output)
	scrollContainer.Direction = container.ScrollHorizontalOnly
	c.scroll = scrollContainer

	c.window = app.NewWindow("GoCalc")

	header := container.NewGridWithColumns(4,
		c.warningButton("C", c.clear), c.strButton("π"), c.strButton("e"), c.addButton("⌫", c.backspace),
	)
	mathBlock := container.NewGridWithRows(2,
		container.NewGridWithColumns(4, c.strButton("sin("), c.strButton("cos("), c.strButton("tan("), c.strButton("√(")),
		container.NewGridWithColumns(4, c.strButton("("), c.strButton(")"), c.strButton("^"), c.strButton("/")),
	)
	mainDigits := container.NewGridWithRows(3,
		container.NewGridWithColumns(4, c.strButton("7"), c.strButton("8"), c.strButton("9"), c.opButton("*")),
		container.NewGridWithColumns(4, c.strButton("4"), c.strButton("5"), c.strButton("6"), c.opButton("-")),
		container.NewGridWithColumns(4, c.strButton("1"), c.strButton("2"), c.strButton("3"), c.opButton("+")),
	)
	footer := container.NewGridWithColumns(2,
		container.NewGridWithColumns(2, c.strButton("."), c.strButton("0")),
		c.eqButton(),
	)

	displayContainer := container.NewThemeOverride(
		container.New(gui.NewTopLayout(),
			container.NewVBox(
				container.NewPadded(c.scroll),
				widget.NewSeparator(),
			),
		),
		gui.NewLargeTextTheme(theme.DefaultTheme(), c.window, c.output),
	)
	buttonsContainer := container.New(gui.NewAdaptiveLayout(7),
		header,
		mathBlock.Objects[0], mathBlock.Objects[1],
		mainDigits.Objects[0], mainDigits.Objects[1], mainDigits.Objects[2],
		footer,
	)

	content := container.NewStack(
		container.NewPadded(displayContainer),
		container.NewPadded(buttonsContainer),
	)

	c.setupEvents()
	c.window.SetContent(content)
	c.window.Resize(fyne.NewSize(screenWeight, screenHeight))
	c.window.Show()
}

func (c *Calculator) addButton(label string, tapped func()) *widget.Button {
	button := widget.NewButton(label, tapped)
	c.buttons[label] = button
	return button
}

func (c *Calculator) strButton(label string) *widget.Button {
	return c.addButton(label, func() {
		c.display(c.expression + label)
	})
}

func (c *Calculator) warningButton(label string, tapped func()) *widget.Button {
	button := c.addButton(label, tapped)
	button.Importance = widget.WarningImportance
	return button
}

func (c *Calculator) opButton(op string) *widget.Button {
	button := c.strButton(op)
	button.Importance = widget.MediumImportance
	return button
}

func (c *Calculator) eqButton() *widget.Button {
	button := c.addButton("=", c.evaluate)
	button.Importance = widget.HighImportance
	return button
}
