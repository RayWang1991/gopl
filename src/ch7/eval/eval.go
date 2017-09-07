package eval

import (
	"fmt"
	"math"
)

// A Var type identifiers a variable, e.g.,x.
type Var string

// A literal is a numeric constant, e.g.3.131.
type literal float64

// A unary represents a unary operator expression, e.g. -x
type unary struct {
	op rune // + -
	x  Expr
}

// A binary represents a binary operator expression, e.g. x+y.
type binary struct {
	op   rune // + - * /
	x, y Expr
}

// A call represents a function call, e.g. sin(x).
type call struct {
	fn   string // pow, sin, cos ,sqrt
	args []Expr
}

// Env is the environment mapping variables to value
type Env map[string]float64

type Expr interface {
	// Eval returns the value of this Expr in the environment env
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String()string
}

func (v Var) Eval(env Env) float64 {
	return env[string(v)]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (i literal) Eval(env Env) float64 {
	return float64(i)
}

func (i literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported operatoer %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	switch u.op {
	case '+', '-':
		u.x.Check(vars)
	}
	return fmt.Errorf("unsupported unary op %q", u.op)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		d := b.y.Eval(env)
		if d == 0 {
			panic(fmt.Sprintf("divide zero"))
		}
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported operatoer %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	switch b.op {
	// TODO, Eval may check the result
	case '+', '-', '*':
		if err := b.x.Check(vars); err != nil {
			return err
		}
		if err := b.y.Check(vars); err != nil {
			return err
		}
		return nil
	case '/':
		if err := b.x.Check(vars); err != nil {
			return err
		}
		if err := b.y.Check(vars); err != nil {
			return err
		}
	}
	return fmt.Errorf("unsupported binary op %q", b.op)
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "cos":
		return math.Cos(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function %s", c.fn))
}

var numParams = map[string]int{
	"pow":  2,
	"sin":  1,
	"cos":  1,
	"sqrt": 1,
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unsupported function call %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s() has %d args", c.fn, arity)
	}
	for _, a := range c.args {
		if err := a.Check(vars); err != nil {
			return err
		}
	}
	return nil
}
