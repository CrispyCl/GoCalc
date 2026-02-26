package calc

import (
	"log"
	"math"
	"strconv"
	"strings"

	"gocalc/internal/eval"
)

func (c *Calculator) display(expr []string) {
	c.expression = expr

	text := strings.Join(expr, "")
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
	rawExpression := strings.Join(c.expression, "")

	if rawExpression == "" || rawExpression == "error" {
		return
	}

	result, err := eval.Evaluate(rawExpression)
	if err != nil {
		log.Println("Error in calculation", err)
		c.display([]string{"error"})
		return
	}

	rounded := math.Round(result*math.Pow(10, 10)) / math.Pow(10, 10)
	resStr := strconv.FormatFloat(rounded, 'g', 10, 64)

	c.display([]string{resStr})
}

func (c *Calculator) clear() {
	c.display([]string{})
}

func (c *Calculator) backspace() {
	if len(c.expression) == 0 {
		c.clear()
		return
	}

	newExpr := c.expression[:len(c.expression)-1]
	c.display(newExpr)
}
