package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Println(i, v, sum)
		if i%5 != 0 {
			c <- sum
		} // send sum to c
	}
}

func main() {
	s := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y)
}
