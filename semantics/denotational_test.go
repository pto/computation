package semantics

import (
	"fmt"
)

// Demonstrate Add to Go.
func ExampleAdd_Go() {
	fmt.Println(Add{Number{1}, Number{2}}.Go())
	// Output:
	// func(e map[string]string) string { l, _ := strconv.Atoi(func(_ map[string]string) string {return "1"}(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "2"}(e)); return fmt.Sprint(l + r) }
}

// Demonstrate Multiply to Go.
func ExampleMultiply_Go() {
	fmt.Println(Multiply{Variable{"x"}, Add{Number{2}, Number{40}}}.Go())
	// Output:
	// func(e map[string]string) string { l, _ := strconv.Atoi(func(e map[string]string) string { return e["x"] }(e)); r, _ := strconv.Atoi(func(e map[string]string) string { l, _ := strconv.Atoi(func(_ map[string]string) string {return "2"}(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "40"}(e)); return fmt.Sprint(l + r) }(e)); return fmt.Sprint(l * r) }
}

// Demonstrate Assign to Go.
func ExampleAssign_Go() {
	fmt.Println(Assign{Variable{"x"}, Number{42}}.Go())
	// Output:
	// func(e map[string]string) map[string]string { e["x"] = func(_ map[string]string) string {return "42"}(e); return e }
}

// Demonstrate a program.
func ExampleProgram() {
	fmt.Println(Program{Sequence{Assign{Variable{"x"}, Number{1}},
		While{LessThan{Variable{"x"}, Number{5}}, Assign{Variable{"x"},
			Multiply{Variable{"x"}, Number{3}}}}}}.Go())
	// Output:
	// package main; import ("fmt"; "strconv"); func main() { e := map[string]string{}; fmt.Println(func(e map[string]string) map[string]string { return func(e map[string]string) map[string]string { b, _ := strconv.ParseBool(func(e map[string]string) string { l, _ := strconv.Atoi(func(e map[string]string) string { return e["x"] }(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "5"}(e)); return fmt.Sprint(l < r) }(e)); for b { e = func(e map[string]string) map[string]string { e["x"] = func(e map[string]string) string { l, _ := strconv.Atoi(func(e map[string]string) string { return e["x"] }(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "3"}(e)); return fmt.Sprint(l * r) }(e); return e }(e); b, _ = strconv.ParseBool(func(e map[string]string) string { l, _ := strconv.Atoi(func(e map[string]string) string { return e["x"] }(e)); r, _ := strconv.Atoi(func(_ map[string]string) string {return "5"}(e)); return fmt.Sprint(l < r) }(e)) } ; return e }(func(e map[string]string) map[string]string { e["x"] = func(_ map[string]string) string {return "1"}(e); return e }(e)) }(e)) }
}
