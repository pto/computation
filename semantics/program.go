package semantics

type Program struct {
	Stmt Statement
}

func (p Program) Go() string {
	return "package main; import (\"fmt\"; \"strconv\"); " +
		"func main() { e := map[string]string{}; fmt.Println(" +
		p.Stmt.Go() + "(e)) }"
}
