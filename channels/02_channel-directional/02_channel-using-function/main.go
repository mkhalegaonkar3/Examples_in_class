package main

import "fmt"

func main() {
	ch := make(chan int)

	go bar(ch)
	foo(ch)
	//fmt.Println("from main", <-ch)
	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	c <- 45
}

//rcv
func bar(c <-chan int) {
	fmt.Println("from bar", <-c)
}
