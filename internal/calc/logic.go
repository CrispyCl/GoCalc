package calc

import (
	"log"
	"math"
	"strconv"
	"strings"

	"gocalc/internal/eval"
)

func (c *Calculator) display(text string) {
	c.expression = text
	c.output.SetText(text)

	if c.scroll != nil {
		c.scroll.Offset.X = c.scroll.Content.MinSize().Width - c.scroll.Size().Width
		if c.scroll.Offset.X < 0 {
			c.scroll.Offset.X = 0
		}
		c.scroll.Refresh()
	}

	if c.window != nil && c.window.Content() != nil {
		c.window.Content().Refresh()
	}
}

func (c *Calculator) evaluate() {
	if c.expression == "" || c.expression == "error" {
		return
	}

	if strings.Contains(c.expression, "error") {
		c.display("error")
		return
	}

	result, err := eval.Evaluate(c.expression)
	if err != nil {
		log.Println("Error in calculation", err)
		c.display("error")
		return
	}

	rounded := math.Round(result*math.Pow(10, 10)) / math.Pow(10, 10)

	c.display(strconv.FormatFloat(rounded, 'g', 10, 64))
}

func (c *Calculator) clear() {
	c.display("")
}

func (c *Calculator) backspace() {
	if c.expression == "" || c.expression == "error" {
		c.clear()
		return
	}

	runes := []rune(c.expression)
	c.display(string(runes[:len(runes)-1]))
}
