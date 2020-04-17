package main

import "fmt"

type t struct {
	a int
}

func New(i int) *t {
	return &t{a: i}
}

func (m *t) Change(i int) {
	m.a = i
}

func (m *t) Print(s string) {
	fmt.Printf("%s = %d\n", s, m.a)
}

func main() {
	a := New(1)
	a.Print("a")

	// copy by this
	b := *a
	a.Print("a")
	b.Print("b")

	a.Change(3)
	a.Print("a")
	b.Print("b")
}
