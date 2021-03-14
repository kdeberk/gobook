package eval

import (
	"fmt"
	"math"
	"strings"
)

type VarSet map[Var]bool

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(VarSet) error
	// String returns a pretty String representation of the Expr
	String() string
}

// A Var identifies a variable, e.g., x.
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vs VarSet) error {
	vs[v] = true
	return nil
}

func (v Var) String() string {
	return string(v)
}

// A literal is a numeric constant, e.g., 3.141.
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (literal) Check(VarSet) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%v", float64(l))
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
	}
}

func (u unary) Check(vs VarSet) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary operator: %q", u.op)
	}
	return u.x.Check(vs)
}

func (u unary) String() string {
	return fmt.Sprintf("%v%v", string(u.op), u.x.String())
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	x, y := b.x.Eval(env), b.y.Eval(env)

	switch b.op {
	case '+':
		return x + y
	case '-':
		return x - y
	case '*':
		return x * y
	case '/':
		return x / y
	default:
		panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
	}
}

func (b binary) Check(vs VarSet) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected unary operator: %q", b.op)
	}
	if err := b.x.Check(vs); err != nil {
		return err
	}
	return b.y.Check(vs)
}

func (u binary) String() string {
	lexpr := u.x.String()
	rexpr := u.y.String()

	if strings.ContainsRune("*/", u.op) {
		if x, ok := u.x.(binary); ok && strings.ContainsRune("-+", x.op) {
			lexpr = "(" + lexpr + ")"
		}
		if y, ok := u.y.(binary); ok && strings.ContainsRune("-+", y.op) {
			rexpr = "(" + rexpr + ")"
		}
	}
	return fmt.Sprintf("%v %v %v", lexpr, string(u.op), rexpr)
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Sprintf("unsupported function call: %q", c.fn))
	}
}

var numParams = map[string]int{"pow": 2, "sqrt": 1, "sin": 1}

func (c call) Check(vs VarSet) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unsupported function call: %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vs); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	var as []string
	for _, a := range c.args {
		as = append(as, a.String())
	}

	return fmt.Sprintf("%v(%v)", c.fn, strings.Join(as, ", "))
}

type Env map[Var]float64
