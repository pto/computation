package small

import "os"

var e Environment = make(map[Variable]Expression)

// Demonstrate reduction of Add and Multiply.
func ExampleAdd() {
	m := &ExpressionMachine{Add{
		Multiply{Number{1}, Number{2}},
		Multiply{Number{3}, Number{4}}}, e}
	m.Run(os.Stdout)
	// Output:
	// 1 * 2 + 3 * 4
	// 2 + 3 * 4
	// 2 + 12
	// 14
}

// Demonstrate reduction of LessThan.
func ExampleBoolean() {
	m := &ExpressionMachine{LessThan{
		Number{5}, Add{Number{2}, Number{2}}}, e}
	m.Run(os.Stdout)
	// Output:
	// 5 < 2 + 2
	// 5 < 4
	// false
}

// Demonstrates an assignment.
func ExampleAssign() {
	e := Environment{Variable{"x"}: Number{1}}
	m := &Machine{Assign{Variable{"x"}, Add{Variable{"x"}, Number{42}}}, e}
	m.Run(os.Stdout)
	// Output:
	// x = x + 42 map[x:1]
	// x = 1 + 42 map[x:1]
	// x = 43 map[x:1]
	// do-nothing map[x:43]
}
