package calc

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *Calculator) LoadUI(app fyne.App) {
	c.output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.output.TextStyle = fyne.TextStyle{Monospace: true}

	c.window = app.NewWindow("GoCalc")

	header := container.NewGridWithColumns(4,
		c.warningButton("C", c.clear),
		c.strButton("π"),
		c.strButton("e"),
		c.addButton("⌫", c.backspace),
	)

	mathBlock := container.NewGridWithColumns(4,
		c.strButton("sin("),
		c.strButton("cos("),
		c.strButton("tan("),
		c.strButton("√("),
		c.strButton("("),
		c.strButton(")"),
		c.strButton("^"),
		c.strButton("/"),
	)

	mainDigits := container.NewGridWithColumns(4,
		c.strButton("7"), c.strButton("8"), c.strButton("9"), c.opButton("*"),
		c.strButton("4"), c.strButton("5"), c.strButton("6"), c.opButton("-"),
		c.strButton("1"), c.strButton("2"), c.strButton("3"), c.opButton("+"),
	)

	footer := container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
			c.strButton("."), c.strButton("0"),
		),
		c.eqButton(),
	)

	buttonsContainer := container.NewVBox(
		header,
		mathBlock,
		mainDigits,
		footer,
	)

	c.window.SetContent(container.NewBorder(
		container.NewVBox(container.NewPadded(c.output), widget.NewSeparator()),
		buttonsContainer,
		nil, nil,
		nil,
	))

	c.setupEvents()
	c.window.Resize(fyne.NewSize(300, 300))
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
