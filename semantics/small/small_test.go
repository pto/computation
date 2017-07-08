package small

import "os"

var e Environment = make(map[Variable]Expression)

// Demonstrate reduction.
func ExampleReduction() {
	m := &Machine{Add{
		Multiply{Number{1}, Number{2}},
		Multiply{Number{3}, Number{4}}}, e}
	m.Run(os.Stdout)
	// Output:
	// 1 * 2 + 3 * 4
	// 2 + 3 * 4
	// 2 + 12
	// 14
}

// Demonstrate booleans.
func ExampleBoolean() {
	m := &Machine{LessThan{
		Number{5}, Add{Number{2}, Number{2}}}, e}
	m.Run(os.Stdout)
	// Output:
	// 5 < 2 + 2
	// 5 < 4
	// false
}
