package small

import "fmt"

// showReduction shows the steps to reduce an expression.
func showReduction(e Expression) {
	for {
		fmt.Println(e)
		if !e.Reducible() {
			fmt.Println("Not Reducible")
			break
		}
		e = e.Reduce()
	}
}

// Demonstrate reduction
func ExampleReduction() {
	var e Expression = Add{
		Multiply{Number{1}, Number{2}},
		Multiply{Number{3}, Number{4}}}
	showReduction(e)
	// Output:
	// 1 * 2 + 3 * 4
	// 2 + 3 * 4
	// 2 + 12
	// 14
	// Not Reducible
}

func ExampleBoolean() {
	var e Expression = LessThan{
		Number{2}, Number{1}}
	showReduction(e)
	// Output:
	// 2 < 1
	// false
	// Not Reducible
}
