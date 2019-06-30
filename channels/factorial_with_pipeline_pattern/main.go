package main

import "fmt"

func main() {
	c := gen()
	out := factorial(c)

	for n := range out {
		fmt.Println(n)
	}
}
func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}

		}
		close(out)
	}()
	return out
}

func factorial(n chan int) chan int {
	out := make(chan int)

	go func() {
		for in := range n {
			out <- fact(in)
		}

		close(out)
	}()

	return out
}
func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i

	}

	return total
}
