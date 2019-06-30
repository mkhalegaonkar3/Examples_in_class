package main

import "fmt"

func main() {
	c := make(chan int)

	send := make(chan<- int)
	rcv := make(<-chan int)

	send <- 45
	<-rcv

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", send)
	fmt.Printf("%T\n", rcv)

}
