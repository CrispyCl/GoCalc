package calc

import (
	"strconv"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func setupCalc() *Calculator {
	c := NewCalculator()
	c.LoadUI(test.NewApp())
	return c
}

func TestCalculator_Arithmetic(t *testing.T) {
	testCases := []struct {
		name     string
		taps     []string
		expected string
	}{
		{"Addition", []string{"1", "+", "1", "="}, "2"},
		{"Substraction", []string{"2", "-", "3", "="}, "-1"},
		{"Multiplication", []string{"5", "*", "2", "="}, "10"},
		{"Division", []string{"3", "/", "2", "="}, "1.5"},
		{"Parenthesis", []string{"2", "*", "(", "3", "+", "4", ")", "="}, "14"},
		{"Dot", []string{"2", ".", "2", "+", "7", ".", "8", "="}, "10"},
		{"Clear", []string{"1", "2", "C"}, ""},
	}

	for _, tс := range testCases {
		t.Run(tс.name, func(t *testing.T) {
			c := setupCalc()
			for _, label := range tс.taps {
				if btn, ok := c.buttons[label]; ok {
					test.Tap(btn)
				} else {
					t.Fatalf("button %s not found", label)
				}
			}
			assert.Equal(t, tс.expected, c.output.Text)
		})
	}
}

func TestCalculator_Scientific(t *testing.T) {
	c := setupCalc()

	testCases := []struct {
		name     string
		taps     []string
		expected string
		delta    float64
	}{
		{
			name:     "Square Root",
			taps:     []string{"√(", "1", "6", ")", "="},
			expected: "4",
		},
		{
			name:     "Power",
			taps:     []string{"2", "^", "3", "="},
			expected: "8",
		},
		{
			name:     "Sin",
			taps:     []string{"sin(", "π", "/", "2", ")", "="},
			expected: "1",
			delta:    1e-9,
		},
		{
			name:     "Cos",
			taps:     []string{"cos(", "π", ")", "="},
			expected: "-1",
			delta:    1e-9,
		},
		{
			name:     "Tan",
			taps:     []string{"tan(", "π", "/", "4", ")", "="},
			expected: "1",
			delta:    1e-9,
		},
		{
			name:     "Constant E",
			taps:     []string{"e", "="},
			expected: "2.718281828",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			test.Tap(c.buttons["C"])

			for _, label := range tc.taps {
				btn, ok := c.buttons[label]
				if !ok {
					t.Fatalf("button %s not found", label)
				}
				test.Tap(btn)
			}

			if tc.delta > 0 {
				actual, _ := strconv.ParseFloat(c.output.Text, 64)
				exp, _ := strconv.ParseFloat(tc.expected, 64)
				assert.InDelta(t, exp, actual, tc.delta)
			} else {
				assert.Equal(t, tc.expected, c.output.Text)
			}
		})
	}
}

func TestCalculator_Keyboard(t *testing.T) {
	c := setupCalc()

	test.TypeOnCanvas(c.window.Canvas(), "1+1")
	assert.Equal(t, "1+1", c.output.Text)

	c.onTypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	assert.Equal(t, "2", c.output.Text)

	test.TypeOnCanvas(c.window.Canvas(), "c")
	assert.Equal(t, "", c.output.Text)
}

func TestCalculator_Backspace(t *testing.T) {
	c := setupCalc()

	c.onTypedKey(&fyne.KeyEvent{Name: fyne.KeyBackspace})
	assert.Equal(t, "", c.output.Text)

	test.TypeOnCanvas(c.window.Canvas(), "1/2")
	c.onTypedKey(&fyne.KeyEvent{Name: fyne.KeyBackspace})
	assert.Equal(t, "1/", c.output.Text)

	c.onTypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	assert.Equal(t, "error", c.output.Text)

	c.onTypedKey(&fyne.KeyEvent{Name: fyne.KeyBackspace})
	assert.Equal(t, "", c.output.Text)
}

func TestCalculator_Shortcuts(t *testing.T) {
	app := test.NewApp()
	c := NewCalculator()
	c.LoadUI(app)
	clipboard := app.Clipboard()

	// Test Copy
	c.display("720+80")
	c.onCopyShortcut(&fyne.ShortcutCopy{Clipboard: clipboard})
	assert.Equal(t, "720+80", clipboard.Content())

	// Test Paste
	c.clear()
	clipboard.SetContent("850")
	c.onPasteShortcut(&fyne.ShortcutPaste{Clipboard: clipboard})
	assert.Equal(t, "850", c.output.Text)

	// Paste invalid data
	clipboard.SetContent("not-a-number")
	c.onPasteShortcut(&fyne.ShortcutPaste{Clipboard: clipboard})
	assert.Equal(t, "850", c.output.Text, "Should not paste non-numeric strings")
}

func TestCalculator_Error(t *testing.T) {
	c := setupCalc()

	test.TypeOnCanvas(c.window.Canvas(), "1/=")
	assert.Equal(t, "error", c.output.Text)
	test.TypeOnCanvas(c.window.Canvas(), "c")

	test.TypeOnCanvas(c.window.Canvas(), "()13=")
	assert.Equal(t, "error", c.output.Text)
}
