package eval

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate_Arithmetic(t *testing.T) {
	testCases := []struct {
		name     string
		expr     string
		expected float64
	}{
		{"Addition", "2 + 5", 7},
		{"Subtraction", "13 - 2", 11},
		{"Multiplication", "25 * 3", 75},
		{"Division", "55 / 2", 27.5},
		{"Priority Of Operations", "10 - 2 * 3", 4},
		{"Parentheses", "(10 - 2) * 3", 24},
		{"Type Conversion", "10 / 2", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Evaluate(tc.expr)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
			assert.IsType(t, float64(0), res)
		})
	}
}

func TestEvaluate_Scientific(t *testing.T) {
	testCases := []struct {
		name     string
		expr     string
		expected float64
		delta    float64 // For floating point comparison
	}{
		{"Constant PI (text)", "pi", math.Pi, 0},
		{"Constant PI (symbol)", "π", math.Pi, 0},
		{"Constant E", "e", math.E, 0},
		{"Square Root", "√(16)", 4, 0},
		{"Power", "3 ^ 3", 27, 0},
		{"Sin(π/2)", "sin(pi * 1/2)", 1.0, 1e-9},
		{"Cos(π/2)", "cos(pi * 1/2)", 0.0, 1e-9},
		{"Tan(π/4)", "tan(pi / 4)", 1.0, 1e-9},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Evaluate(tc.expr)
			assert.NoError(t, err)
			if tc.delta > 0 {
				assert.InDelta(t, tc.expected, res, tc.delta)
			} else {
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}

func TestEvaluate_ValidationErrors(t *testing.T) {
	testCases := []struct {
		name     string
		expr     string
		expected error
	}{
		{"Syntax Error", "5 + * 2", nil}, // nil means we just check for generic error
		{"Logical Operator", "10 > 5", ErrUnexpectedReturnType},
		{"Division by Zero", "10 / 0", ErrMath},
		{"Empty String", "", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Evaluate(tc.expr)
			assert.Error(t, err)
			if tc.expected != nil {
				assert.ErrorIs(t, err, tc.expected)
			}
		})
	}
}
