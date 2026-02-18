package eval

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate_BasicArithmetic(t *testing.T) {
	// Addition
	res, err := Evaluate("2 + 5")
	assert.NoError(t, err)
	assert.Equal(t, float64(7), res)

	// Subtraction
	res, err = Evaluate("13 - 2")
	assert.NoError(t, err)
	assert.Equal(t, float64(11), res)

	// Multiplication
	res, err = Evaluate("25 * 3")
	assert.NoError(t, err)
	assert.Equal(t, float64(75), res)

	// division
	res, err = Evaluate("55 / 2")
	assert.NoError(t, err)
	assert.Equal(t, float64(27.5), res)

	// Priority of operations
	res, err = Evaluate("10 - 2 * 3")
	assert.NoError(t, err)
	assert.Equal(t, float64(4), res)

	// Parentheses
	res, err = Evaluate("(10 - 2) * 3")
	assert.NoError(t, err)
	assert.Equal(t, float64(24), res)
}

func TestEvaluate_MathFunctions(t *testing.T) {
	// PI constant
	res, err := Evaluate("pi")
	assert.NoError(t, err)
	assert.Equal(t, math.Pi, res)

	res, err = Evaluate("π")
	assert.NoError(t, err)
	assert.Equal(t, math.Pi, res)

	// E constant
	res, err = Evaluate("e")
	assert.NoError(t, err)
	assert.Equal(t, math.E, res)

	// sqrt
	res, err = Evaluate("√(16)")
	assert.NoError(t, err)
	assert.Equal(t, float64(4), res)

	// pow
	res, err = Evaluate("3 ^ 3")
	assert.NoError(t, err)
	assert.Equal(t, float64(27), res)

	// Sin of Pi/2
	res, err = Evaluate("sin(pi * 1/2)")
	assert.NoError(t, err)
	assert.InDelta(t, 1.0, res, 1e-9)

	// Cos of Pi/2
	res, err = Evaluate("cos(pi * 1/2)")
	assert.NoError(t, err)
	assert.InDelta(t, 0.0, res, 1e-9)

	// Tan of Pi/4
	res, err = Evaluate("tan(pi / 4)")
	assert.NoError(t, err)
	assert.InDelta(t, 1.0, res, 1e-9)
}

func TestEvaluate_Types(t *testing.T) {
	// Checking the conversion of int results to float64
	res, err := Evaluate("10 / 2")
	assert.NoError(t, err)
	assert.IsType(t, float64(0), res)
	assert.Equal(t, float64(5), res)
}

func TestEvaluate_Errors(t *testing.T) {
	// Sintax error
	_, err := Evaluate("5 + * 2")
	assert.Error(t, err)

	// Unexpected return type
	_, err = Evaluate("10 > 5")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrUnexpectedReturnType)
}
