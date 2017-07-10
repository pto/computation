package semantics

import (
	"fmt"
)

// Demonstrate Add to Go.
func ExampleAdd_denotational() {
	fmt.Println(Add{Number{1}, Number{2}}.Go())
	// Output:
	// func(e map[string]string) string { l, _ := strconv.Atoi(func(_ map[string]string) string {return "1"}(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "2"}(e)); return fmt.Sprint(l + r) }
}

// Demonstrate Multiply to Go.
func ExampleMultiply_denotational() {
	fmt.Println(Multiply{Variable{"x"}, Add{Number{2}, Number{40}}}.Go())
	// Output:
	// func(e map[string]string) string { l, _ := strconv.Atoi(func(e map[string]string) string { return e["x"] }(e)); r, _ := strconv.Atoi(func(e map[string]string) string { l, _ := strconv.Atoi(func(_ map[string]string) string {return "2"}(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "40"}(e)); return fmt.Sprint(l + r) }(e)); return fmt.Sprint(l * r) }
}
