package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
		var n1=0
		var n2=1
		return func() int{
			var out int
			out = n1
			n1 = n2
			n2 = out + n1
			return out
		}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}