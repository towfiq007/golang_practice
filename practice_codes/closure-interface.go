package main

import "fmt"

type I interface {
	M()
	N()
}

type T struct {
	a string
	b string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.a, t.b)
}
func (t T) N() {
	fmt.Println(t.a, t.b)
}

func main() {
	var i I = T{"aaaa", "bbbb"}
	i.M()
	i.N()
}
