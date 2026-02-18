package calc

import (
	"log"
	"strconv"
	"strings"

	"gocalc/internal/eval"
)

func (c *Calculator) display(text string) {
	c.expression = text
	c.output.SetText(text)
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

	c.display(strconv.FormatFloat(result, 'f', -1, 64))
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
