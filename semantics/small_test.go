package semantics

import "os"

// Demonstrate reduction of Add and Multiply.
func ExampleAdd() {
	e := Environment{}
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
	e := Environment{}
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

// Demonstrates an If statement.
func ExampleIf() {
	e := Environment{Variable{"x"}: Boolean{true}}
	m := &Machine{If{Variable{"x"}, Assign{Variable{"x"}, Number{1}},
		Assign{Variable{"x"}, Number{2}}}, e}
	m.Run(os.Stdout)
	// Output:
	// if x {x = 1} else {x = 2} map[x:true]
	// if true {x = 1} else {x = 2} map[x:true]
	// x = 1 map[x:true]
	// do-nothing map[x:1]
}

// Demonstrates a sequence.
func ExampleSequence() {
	e := Environment{}
	m := &Machine{Sequence{Assign{Variable{"x"}, Add{Number{1}, Number{1}}},
		Assign{Variable{"x"}, Add{Variable{"x"}, Number{3}}}}, e}
	m.Run(os.Stdout)
	// Output:
	// x = 1 + 1; x = x + 3 map[]
	// x = 2; x = x + 3 map[]
	// do-nothing; x = x + 3 map[x:2]
	// x = x + 3 map[x:2]
	// x = 2 + 3 map[x:2]
	// x = 5 map[x:2]
	// do-nothing map[x:5]
}

// Demonstrates a while statement.
func ExampleWhile() {
	e := Environment{Variable{"x"}: Number{1}}
	m := &Machine{While{LessThan{Variable{"x"}, Number{5}},
		Assign{Variable{"x"}, Multiply{Variable{"x"}, Number{3}}}}, e}
	m.Run(os.Stdout)
	// Output:
	// while x < 5 {x = x * 3} map[x:1]
	// if x < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:1]
	// if 1 < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:1]
	// if true {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:1]
	// x = x * 3; while x < 5 {x = x * 3} map[x:1]
	// x = 1 * 3; while x < 5 {x = x * 3} map[x:1]
	// x = 3; while x < 5 {x = x * 3} map[x:1]
	// do-nothing; while x < 5 {x = x * 3} map[x:3]
	// while x < 5 {x = x * 3} map[x:3]
	// if x < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:3]
	// if 3 < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:3]
	// if true {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:3]
	// x = x * 3; while x < 5 {x = x * 3} map[x:3]
	// x = 3 * 3; while x < 5 {x = x * 3} map[x:3]
	// x = 9; while x < 5 {x = x * 3} map[x:3]
	// do-nothing; while x < 5 {x = x * 3} map[x:9]
	// while x < 5 {x = x * 3} map[x:9]
	// if x < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:9]
	// if 9 < 5 {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:9]
	// if false {x = x * 3; while x < 5 {x = x * 3}} else {do-nothing} map[x:9]
	// do-nothing map[x:9]
}
