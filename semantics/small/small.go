// Package small implements a simple AST language using small-step semantics.
package small

import "fmt"

// Expressions are a combination of numbers, additions and multiplications.
type Expression interface {
	String() string
	Reducible() bool
	Reduce() Expression
}

// Number represents an integer.
type Number struct {
	Value int
}

func (n Number) String() string {
	return fmt.Sprint(n.Value)
}

func (n Number) Reducible() bool {
	return false
}

func (n Number) Reduce() Expression {
	panic("Cannot reduce a Number")
}

// Boolean represents a boolean value.
type Boolean struct {
	Value bool
}

func (b Boolean) String() string {
	return fmt.Sprint(b.Value)
}

func (b Boolean) Reducible() bool {
	return false
}

func (b Boolean) Reduce() Expression {
	panic("Cannot reduce a Boolean")
}

// Add represents an addition of two expressions.
type Add struct {
	Left  Expression
	Right Expression
}

func (a Add) String() string {
	return a.Left.String() + " + " + a.Right.String()
}

func (a Add) Reducible() bool {
	return true
}

func (a Add) Reduce() Expression {
	if a.Left.Reducible() {
		return Add{a.Left.Reduce(), a.Right}
	} else if a.Right.Reducible() {
		return Add{a.Left, a.Right.Reduce()}
	} else {
		return Number{a.Left.(Number).Value + a.Right.(Number).Value}
	}
}

// Multiply represents a multiplication of two expressions.
type Multiply struct {
	Left  Expression
	Right Expression
}

func (m Multiply) String() string {
	return m.Left.String() + " * " + m.Right.String()
}

func (m Multiply) Reducible() bool {
	return true
}

func (m Multiply) Reduce() Expression {
	if m.Left.Reducible() {
		return Multiply{m.Left.Reduce(), m.Right}
	} else if m.Right.Reducible() {
		return Multiply{m.Left, m.Right.Reduce()}
	} else {
		return Number{m.Left.(Number).Value * m.Right.(Number).Value}
	}
}

// LessThan represents the comparision of two expressions.
type LessThan struct {
	Left  Expression
	Right Expression
}

func (lt LessThan) String() string {
	return lt.Left.String() + " < " + lt.Right.String()
}

func (lt LessThan) Reducible() bool {
	return true
}

func (lt LessThan) Reduce() Expression {
	if lt.Left.Reducible() {
		return LessThan{lt.Left.Reduce(), lt.Right}
	} else if lt.Right.Reducible() {
		return LessThan{lt.Left, lt.Right.Reduce()}
	} else {
		switch lt.Left.(type) {
		case Number:
			return Boolean{lt.Left.(Number).Value < lt.Right.(Number).Value}
		case Boolean:
			panic("Cannot compare Booleans")
		default:
			panic("Unknown type")
		}
	}
}
