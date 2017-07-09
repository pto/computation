// Package semantics implements a simple AST language using either small-step
// or big-step semantics.
package semantics

import (
	"fmt"
)

// Environment maps variables to expressions.
type Environment map[Variable]Expression

// Expression is a combination of numbers, additions and multiplications.
type Expression interface {
	String() string
	Reducible() bool
	Reduce(Environment) Expression
	Evaluate(Environment) Expression
}

// Number represents an integer.
type Number struct {
	Value int
}

// String formats the Number's value.
func (n Number) String() string {
	return fmt.Sprint(n.Value)
}

// Reducible always returns false for a Number.
func (n Number) Reducible() bool {
	return false
}

// Reduce cannot be called on a number.
func (n Number) Reduce(_ Environment) Expression {
	panic("Cannot reduce a Number")
}

// Evaluate on a number returns itself.
func (n Number) Evaluate(_ Environment) Expression {
	return n
}

// Boolean represents a boolean value.
type Boolean struct {
	Value bool
}

// String formats a Boolean as "true" or "false".
func (b Boolean) String() string {
	return fmt.Sprint(b.Value)
}

// Reducible is always false for a Boolean.
func (b Boolean) Reducible() bool {
	return false
}

// Reduce cannot be called on a boolean.
func (b Boolean) Reduce(_ Environment) Expression {
	panic("Cannot reduce a boolean")
}

// Evaluate on a boolean returns itself.
func (b Boolean) Evaluate(_ Environment) Expression {
	return b
}

// Add represents an addition of two expressions.
type Add struct {
	Left  Expression
	Right Expression
}

// String formats the two expressions and joins them with "+".
func (a Add) String() string {
	return a.Left.String() + " + " + a.Right.String()
}

// Reducible is always true for Add.
func (a Add) Reducible() bool {
	return true
}

// Reduce on an Add reduces the left or right sides, or performs the addition.
func (a Add) Reduce(e Environment) Expression {
	if a.Left.Reducible() {
		return Add{a.Left.Reduce(e), a.Right}
	} else if a.Right.Reducible() {
		return Add{a.Left, a.Right.Reduce(e)}
	} else {
		return Number{a.Left.(Number).Value + a.Right.(Number).Value}
	}
}

// Evaluate on an Add returns a number
func (a Add) Evaluate(e Environment) Expression {
	return Number{a.Left.Evaluate(e).(Number).Value +
		a.Right.Evaluate(e).(Number).Value}
}

// Multiply represents a multiplication of two expressions.
type Multiply struct {
	Left  Expression
	Right Expression
}

// String formats the left and right expressions, joined by "*".
func (m Multiply) String() string {
	return m.Left.String() + " * " + m.Right.String()
}

// Reducible is always true for Multiply.
func (m Multiply) Reducible() bool {
	return true
}

// Reduce on a Multiply reduces the left or right expression, or performs the
// multiplication.
func (m Multiply) Reduce(e Environment) Expression {
	if m.Left.Reducible() {
		return Multiply{m.Left.Reduce(e), m.Right}
	} else if m.Right.Reducible() {
		return Multiply{m.Left, m.Right.Reduce(e)}
	} else {
		return Number{m.Left.(Number).Value * m.Right.(Number).Value}
	}
}

// Evaluate on a Multiply returns a number
func (m Multiply) Evaluate(e Environment) Expression {
	return Number{m.Left.Evaluate(e).(Number).Value *
		m.Right.Evaluate(e).(Number).Value}
}

// LessThan represents the comparision of two expressions.
type LessThan struct {
	Left  Expression
	Right Expression
}

// String formats the left and right expressions, joined by "<".
func (lt LessThan) String() string {
	return lt.Left.String() + " < " + lt.Right.String()
}

// Reducible is always true for a LessThan.
func (lt LessThan) Reducible() bool {
	return true
}

// Reduce on a LessThan reduces the left or right expression, or performs
// the less than comparison.
func (lt LessThan) Reduce(e Environment) Expression {
	if lt.Left.Reducible() {
		return LessThan{lt.Left.Reduce(e), lt.Right}
	} else if lt.Right.Reducible() {
		return LessThan{lt.Left, lt.Right.Reduce(e)}
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

// Evaluate on a LessThan returns a boolean
func (l LessThan) Evaluate(e Environment) Expression {
	return Boolean{l.Left.Evaluate(e).(Number).Value <
		l.Right.Evaluate(e).(Number).Value}
}

// Variable represents a name referring to an expression.
type Variable struct {
	Name string
}

// String returns the Variable's name.
func (v Variable) String() string {
	return v.Name
}

// Reducible is always true for a Variable.
func (v Variable) Reducible() bool {
	return true
}

// Reduce on a Variable looks up the Variable's expression.
func (v Variable) Reduce(e Environment) Expression {
	return e[v]
}

// Evaluate on a Variable returns the Variable's expression.
func (v Variable) Evaluate(e Environment) Expression {
	return e[v]
}
