package small

// Statement represents an instruction that could update the environment.
type Statement interface {
	String() string
	Reducible() bool
	Reduce(Environment) (Statement, Environment)
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
	e[a.Var] = a.Exp.Reduce(e)
	return DoNothing{}, e
}
