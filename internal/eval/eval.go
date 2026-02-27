package eval

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/expr-lang/expr"
)

var (
	ErrUnexpectedReturnType = errors.New("unexpected return type")
	ErrMath                 = errors.New("math error")
)

func Evaluate(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, "√", "sqrt")
	expression = strings.ReplaceAll(expression, "^", "**")
	expression = strings.ReplaceAll(expression, "π", "pi")

	env := map[string]any{
		"pi":   math.Pi,
		"e":    math.E,
		"sin":  math.Sin,
		"cos":  math.Cos,
		"tan":  math.Tan,
		"sqrt": math.Sqrt,
	}

	program, err := expr.Compile(expression, expr.Env(env))
	if err != nil {
		return 0, fmt.Errorf("failed to compile expression %q: %w", expression, err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		return 0, fmt.Errorf("failed to run expression: %w", err)
	}

	var result float64
	switch v := output.(type) {
	case float64:
		result = v
	case int:
		result = float64(v)
	case int64:
		result = float64(v)
	case float32:
		result = float64(v)
	default:
		return 0, fmt.Errorf("%w: got %T (value %v)", ErrUnexpectedReturnType, v, v)
	}

	if math.IsInf(result, 0) || math.IsNaN(result) {
		return 0, fmt.Errorf("%w: invalid result (inf or nan)", ErrMath)
	}

	return result, nil
}
