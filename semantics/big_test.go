package semantics

import (
	"fmt"
)

// Demonstrate Add and Multiply.
func ExampleAdd_big() {
	e := Environment{}
	fmt.Println(Assign{Variable{"result"}, Add{
		Multiply{Number{1}, Number{2}},
		Multiply{Number{3}, Number{4}}}}.Evaluate(e))
	// Output:
	// map[result:14]
}

// Demonstrate reduction of LessThan.
func ExampleBoolean_big() {
	e := Environment{}
	fmt.Println(Assign{Variable{"result"}, LessThan{
		Number{5}, Add{Number{2}, Number{2}}}}.Evaluate(e))
	// Output:
	// map[result:false]
}

// Demonstrates an assignment.
func ExampleAssign_big() {
	e := Environment{Variable{"x"}: Number{1}}
	fmt.Println(Assign{Variable{"x"}, Add{Variable{"x"},
		Number{42}}}.Evaluate(e))
	// Output:
	// map[x:43]
}

// Demonstrates an If statement.
func ExampleIf_big() {
	e := Environment{Variable{"x"}: Boolean{true}}
	fmt.Println(If{Variable{"x"}, Assign{Variable{"x"}, Number{1}},
		Assign{Variable{"x"}, Number{2}}}.Evaluate(e))
	// Output:
	// map[x:1]
}

// Demonstrates a sequence.
func ExampleSequence_big() {
	e := Environment{}
	fmt.Println(Sequence{Assign{Variable{"x"}, Add{Number{1}, Number{1}}},
		Assign{Variable{"x"}, Add{Variable{"x"}, Number{3}}}}.Evaluate(e))
	// Output:
	// map[x:5]
}

// Demonstrates a while statement.
func ExampleWhile_big() {
	e := Environment{Variable{"x"}: Number{1}}
	fmt.Println(While{LessThan{Variable{"x"}, Number{5}},
		Assign{Variable{"x"}, Multiply{Variable{"x"}, Number{3}}}}.Evaluate(e))
	// Output:
	// map[x:9]
}
