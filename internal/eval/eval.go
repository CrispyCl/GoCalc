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
)

func Evaluate(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, "√", "sqrt")
	expression = strings.ReplaceAll(expression, "^", "**")

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

	switch v := output.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("%w: got %T (value %v)", ErrUnexpectedReturnType, v, v)
	}
}
