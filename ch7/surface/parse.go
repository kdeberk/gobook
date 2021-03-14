package main

import (
	"fmt"

	"dberk.nl/gobook/ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(eval.VarSet)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		switch v {
		case "x", "y", "r":
			continue
		default:
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}
