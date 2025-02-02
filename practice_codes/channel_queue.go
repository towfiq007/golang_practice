package main

import "fmt"

func r1(c chan int) {
	x := <-c
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go r1(c)
	c <- 1
}
