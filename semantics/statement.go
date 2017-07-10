package semantics

// Statement represents an instruction that could update the environment.
type Statement interface {
	String() string
	Reducible() bool
	Reduce(Environment) (Statement, Environment)
	Evaluate(Environment) Environment
}

// DoNothing is the no-op statement
type DoNothing struct {
}

// String displays nothing.
func (d DoNothing) String() string {
	return "do-nothing"
}

// Reducible on a DoNothing always returns false.
func (d DoNothing) Reducible() bool {
	return false
}

// Reduce is not allowed on a DoNothing.
func (d DoNothing) Reduce(_ Environment) (Statement, Environment) {
	panic("Cannot reduce a DoNothing")
}

// Evaluate on a DoNothing returns an unchanged environment.
func (d DoNothing) Evaluate(e Environment) Environment {
	return e
}

// Assign represents an assignment statement.
type Assign struct {
	Var Variable
	Exp Expression
}

// String displays the assignment statement.
func (a Assign) String() string {
	return a.Var.Name + " = " + a.Exp.String()
}

// Reducible is always true for an Assign.
func (a Assign) Reducible() bool {
	return true
}

// Reduce on an Assign makes the assignment and returns a DoNothing.
func (a Assign) Reduce(e Environment) (Statement, Environment) {
	if a.Exp.Reducible() {
		return Assign{a.Var, a.Exp.Reduce(e)}, e
	}
	e[a.Var] = a.Exp
	return DoNothing{}, e
}

// Evaluate on an assignment returns an updated Environment.
func (a Assign) Evaluate(e Environment) Environment {
	e[a.Var] = a.Exp.Evaluate(e)
	return e
}

// If represents a choice.
type If struct {
	Condition   Expression
	Consequence Statement
	Alternative Statement
}

// String displays the if statement
func (i If) String() string {
	return "if " + i.Condition.String() + " {" + i.Consequence.String() +
		"} else {" + i.Alternative.String() + "}"
}

// Reducible is always true for If
func (i If) Reducible() bool {
	return true
}

// Reduce reduces the condition, or replaces it with the result.
func (i If) Reduce(e Environment) (Statement, Environment) {
	if i.Condition.Reducible() {
		return If{i.Condition.Reduce(e), i.Consequence, i.Alternative}, e
	}
	if i.Condition.(Boolean).Value {
		return i.Consequence, e
	}
	return i.Alternative, e
}

// Evaluate on an If returns the evaluation of one of its statements.
func (i If) Evaluate(e Environment) Environment {
	if i.Condition.Evaluate(e).(Boolean).Value {
		return i.Consequence.Evaluate(e)
	}
	return i.Alternative.Evaluate(e)
}

// Sequence reduces two statements consecutively.
type Sequence struct {
	First  Statement
	Second Statement
}

// String displays the statements of the sequence.
func (s Sequence) String() string {
	return s.First.String() + "; " + s.Second.String()
}

// Reducible is always true for a sequence.
func (s Sequence) Reducible() bool {
	return true
}

// Reduce reduces the first statement unless is it DoNothing.
func (s Sequence) Reduce(e Environment) (Statement, Environment) {
	switch s.First.(type) {
	case DoNothing:
		return s.Second, e
	default:
		reducedFirst, newE := s.First.Reduce(e)
		return Sequence{reducedFirst, s.Second}, newE
	}
}

// Evaluate on a Sequence returns the environment after evalutating both
// statements.
func (s Sequence) Evaluate(e Environment) Environment {
	return s.Second.Evaluate(s.First.Evaluate(e))
}

// While repeats a statement while a condition is true.
type While struct {
	Condition Expression
	Body      Statement
}

// String displays the while construct.
func (w While) String() string {
	return "while " + w.Condition.String() + " {" + w.Body.String() + "}"
}

// Reducible is always true for a while statement.
func (w While) Reducible() bool {
	return true
}

// Reduce unrolls one loop of the while statement by turning it into an If.
func (w While) Reduce(e Environment) (Statement, Environment) {
	return If{w.Condition, Sequence{w.Body, w}, DoNothing{}}, e
}

// Evaluate on a While repeats the body based on the condition.
func (w While) Evaluate(e Environment) Environment {
	for w.Condition.Evaluate(e).(Boolean).Value {
		e = w.Body.Evaluate(e)
	}
	return e
}
