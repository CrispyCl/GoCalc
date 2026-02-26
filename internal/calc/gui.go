package calc

import (
	"strings"

	"gocalc/internal/calc/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	screenHeight = float32(380)
	screenWeight = float32(300)
	operators    = "+-*/^"
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
		gui.NewAdaptiveTextTheme(theme.DefaultTheme(), c.window, 0.07, nil),
	)

	buttonsMaxSize := 22
	buttonsContainer := container.NewThemeOverride(
		container.New(gui.NewAdaptiveLayout(7),
			header,
			mathBlock.Objects[0], mathBlock.Objects[1],
			mainDigits.Objects[0], mainDigits.Objects[1], mainDigits.Objects[2],
			footer,
		),
		gui.NewAdaptiveTextTheme(theme.DefaultTheme(), c.window, 0.04, &buttonsMaxSize),
	)

	content := container.NewStack(
		container.NewPadded(displayContainer),
		container.NewPadded(buttonsContainer),
	)

	c.setupEvents()
	c.window.SetContent(content)
	c.window.Resize(fyne.NewSize(screenWeight, screenHeight))
	c.window.Show()

	if c.window.Content() != nil {
		c.window.Content().Refresh()
	}
}

func (c *Calculator) addButton(label string, tapped func()) *widget.Button {
	button := widget.NewButton(label, tapped)
	c.buttons[label] = button
	return button
}

func (c *Calculator) strButton(label string) *widget.Button {
	return c.addButton(label, func() {
		// If "error" is on screen, clear it before any new input
		if len(c.expression) == 1 && c.expression[0] == "error" {
			c.expression = []string{}
		}

		expr := c.expression

		// Decimal Point Validation
		if label == "." {
			if len(expr) == 0 {
				c.display([]string{"0", "."})
				return
			}

			for i := len(expr) - 1; i >= 0; i-- {
				if strings.ContainsAny(expr[i], operators+"()") && !strings.ContainsAny(expr[i], "0123456789") {
					break
				}
				if expr[i] == "." {
					return
				}
			}

			lastToken := expr[len(expr)-1]
			if strings.ContainsAny(lastToken, operators+"(") && !strings.ContainsAny(lastToken, "0123456789") {
				c.display(append(expr, "0."))
				return
			}
		}

		// Mathematical Operators Validation
		if strings.ContainsAny(label, operators) {
			if len(expr) == 0 || len(expr) == 1 && strings.ContainsAny(expr[0], operators) && !strings.ContainsAny(expr[0], "0123456789") {
				if label == "-" {
					c.display([]string{"-"})
				}
				return
			}

			// If the last character is already an operator, replace it with the new one
			lastToken := expr[len(expr)-1]

			if strings.ContainsAny(lastToken, operators) && !strings.ContainsAny(lastToken, "0123456789") {
				expr = expr[:len(expr)-1]
			}
		}

		c.display(append(expr, label))
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
