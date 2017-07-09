package semantics

import (
	"fmt"
	"io"
)

// ExpressionMachine holds an environment and can reduce expressions.
type ExpressionMachine struct {
	Exp Expression
	Env Environment
}

// Step performs a single reduction on the machine's Expression.
func (m *ExpressionMachine) Step() {
	m.Exp = m.Exp.Reduce(m.Env)
}

// Run performs all possible reductions, sending each step to the writer.
func (m *ExpressionMachine) Run(w io.Writer) {
	for m.Exp.Reducible() {
		fmt.Fprintln(w, m.Exp)
		m.Step()
	}
	fmt.Fprintln(w, m.Exp)
}

// Machine holds an environment and can reduce statements.
type Machine struct {
	Stmt Statement
	Env  Environment
}

// Step performs a single reduction on the Machine's Statement.
func (m *Machine) Step() {
	m.Stmt, m.Env = m.Stmt.Reduce(m.Env)
}

// Run performs all possible reductions, sending each step to the writer.
func (m *Machine) Run(w io.Writer) {
	for m.Stmt.Reducible() {
		fmt.Fprintln(w, m.Stmt, m.Env)
		m.Step()
	}
	fmt.Fprintln(w, m.Stmt, m.Env)
}
